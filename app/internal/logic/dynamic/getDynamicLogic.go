package dynamic

import (
	"context"
	"git.154896.xyz/zhifou/collect/collect"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/dynamic/dynamic"
	"git.154896.xyz/zhifou/like/like"
	"git.154896.xyz/zhifou/pkg/tool"
	userRpc "git.154896.xyz/zhifou/user/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/fx"
	"github.com/zeromicro/go-zero/core/mr"
	"net/http"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
}

func NewGetDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetDynamicLogic {
	return &GetDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    r,
	}
}

func (l *GetDynamicLogic) GetDynamic(in *types.GetDynamicReq) (*types.GetDynamicResp, error) {
	// 尝试获取Uid
	uid := tool.TryGetIdFromReq(l.req, l.svcCtx.Config.JwtAuth.AccessSecret)

	// 获取动态
	dynamics, err := l.svcCtx.DynamicRpc.GetDynamic(l.ctx, &dynamic.GetDynamicReq{
		LastId:  in.LastId,
		PageNum: in.PageNum,
	})
	if err != nil {
		return nil, err
	}

	// 转换和注入数控
	var resp types.GetDynamicResp
	fx.From(func(source chan<- any) {
		for _, v := range dynamics.Data {
			source <- v
		}
	}).ForEach(func(item any) {
		post := item.(*dynamic.CommonResp)
		var commentResp types.CommenPostResp
		_ = copier.Copy(&commentResp, item)

		fns := []func(){
			// 注入用户信息
			func() {
				var user types.CommonUserResp
				u, _ := l.svcCtx.UserRpc.GetOne(l.ctx, &userRpc.GetOneReq{Id: post.UserId})
				_ = copier.Copy(&user, u)
				commentResp.User = &user
			},
			//注入点赞数量
			func() {
				num, err := l.svcCtx.LikeRpc.GetNum(l.ctx, &like.GetNumReq{PostId: post.Id})
				if err != nil {
					return
				}
				commentResp.UpvoteCount = num.Num
			},
			// 添加评论数量
			func() {
				num, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &comment.GetCommentCountReq{PostId: post.Id})
				if err != nil {
					return
				}
				commentResp.CommentCount = num.Num
			},
			// 添加收藏数量
			func() {
				num, err := l.svcCtx.CollectRpc.GetNum(l.ctx, &collect.GetNumReq{PostId: post.Id})
				if err != nil {
					return
				}
				commentResp.CollectionCount = num.Num
			},
		}
		// 登录用户还需要注入点赞、收藏状态
		if uid != "" {
			// 登录用户
			fns = append(fns, func() {
				// 注入收藏状态
				status, err := l.svcCtx.CollectRpc.GetStatus(l.ctx, &collect.GetStatusReq{
					PostId: post.Id,
					UserId: uid,
				})
				if err != nil {
					return
				}
				commentResp.CollectStatus = status.Status
			}, func() {
				// 注入点赞状态
				status, err := l.svcCtx.LikeRpc.GetStatus(l.ctx, &like.GetStatusReq{PostId: post.Id, UserId: uid})
				if err != nil {
					return
				}
				commentResp.UpvoteStatus = status.Status
			})
		}
		mr.FinishVoid(fns...)
		resp.List = append(resp.List, &commentResp)

	})

	return &resp, nil
}
