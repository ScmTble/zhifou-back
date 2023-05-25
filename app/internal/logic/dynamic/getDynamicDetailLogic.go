package dynamic

import (
	"context"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"git.154896.xyz/zhifou/collect/collect"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/dynamic/dynamic"
	"git.154896.xyz/zhifou/like/like"
	"git.154896.xyz/zhifou/pkg/tool"
	userRpc "git.154896.xyz/zhifou/user/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDynamicDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
}

func NewGetDynamicDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *GetDynamicDetailLogic {
	return &GetDynamicDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    r,
	}
}

func (l *GetDynamicDetailLogic) GetDynamicDetail(req *types.GetDynamicDetailReq) (*types.GetDynamicDetailResp, error) {
	var resp types.GetDynamicDetailResp

	// 获取动态数据
	post, err := l.svcCtx.DynamicRpc.FindPost(l.ctx, &dynamic.FindOneReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	_ = copier.Copy(&resp, post)

	fns := []func(){
		// 注入用户信息
		func() {
			var user types.CommonUserResp
			u, _ := l.svcCtx.UserRpc.GetOne(l.ctx, &userRpc.GetOneReq{Id: post.UserId})
			_ = copier.Copy(&user, u)
			resp.User = &user
		},
		//注入点赞数量
		func() {
			num, err := l.svcCtx.LikeRpc.GetNum(l.ctx, &like.GetNumReq{PostId: post.Id})
			if err != nil {
				return
			}
			resp.UpvoteCount = num.Num
		},
		// 添加评论数量
		func() {
			num, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &comment.GetCommentCountReq{PostId: post.Id})
			if err != nil {
				return
			}
			resp.CommentCount = num.Num
		},
		// 添加收藏数量
		func() {
			num, err := l.svcCtx.CollectRpc.GetNum(l.ctx, &collect.GetNumReq{PostId: post.Id})
			if err != nil {
				return
			}
			resp.CollectionCount = num.Num
		},
	}
	// 尝试获取uid
	if uid := tool.TryGetIdFromReq(l.req, l.svcCtx.Config.JwtAuth.AccessSecret); uid != "" {
		// 登录用户
		fns = append(fns, func() {
			// 注入收藏状态
			status, err := l.svcCtx.CollectRpc.GetStatus(l.ctx, &collect.GetStatusReq{
				PostId: req.Id,
				UserId: uid,
			})
			if err != nil {
				return
			}
			resp.CollectStatus = status.Status
		}, func() {
			// 注入点赞状态
			status, err := l.svcCtx.LikeRpc.GetStatus(l.ctx, &like.GetStatusReq{PostId: req.Id, UserId: uid})
			if err != nil {
				return
			}
			resp.UpvoteStatus = status.Status
		})
	}
	mr.FinishVoid(fns...)

	return &resp, nil
}
