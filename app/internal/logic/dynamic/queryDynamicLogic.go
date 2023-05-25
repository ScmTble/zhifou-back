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

type QueryDynamicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	req    *http.Request
}

func NewQueryDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *QueryDynamicLogic {
	return &QueryDynamicLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		req:    r,
	}
}

func (l *QueryDynamicLogic) QueryDynamic(req *types.QueryDynamicReq) (resp *types.QueryDynamicResp, err error) {
	// 尝试获取Uid
	uid := tool.TryGetIdFromReq(l.req, l.svcCtx.Config.JwtAuth.AccessSecret)

	// 分页查询动态
	dynamics, err := l.svcCtx.DynamicRpc.QueryTag(l.ctx, &dynamic.QueryTagReq{
		TagId:   req.TagId,
		LastId:  req.LastId,
		PageNum: req.PageNum,
	})

	if len(dynamics.Data) == 0 {
		return nil, nil
	}

	// 数据填充
	var result []*types.CommenPostResp
	fx.From(func(source chan<- any) {
		for _, v := range dynamics.Data {
			source <- v
		}
	}).ForEach(func(item any) {
		commonResp := item.(*dynamic.CommonResp)
		var resp types.CommenPostResp
		_ = copier.Copy(&resp, item)

		fns := []func(){
			// 注入用户信息
			func() {
				var user types.CommonUserResp
				u, _ := l.svcCtx.UserRpc.GetOne(l.ctx, &userRpc.GetOneReq{Id: commonResp.UserId})
				_ = copier.Copy(&user, u)
				resp.User = &user
			},
			//注入点赞数量
			func() {
				num, err := l.svcCtx.LikeRpc.GetNum(l.ctx, &like.GetNumReq{PostId: commonResp.Id})
				if err != nil {
					return
				}
				resp.UpvoteCount = num.Num
			},
			// 添加评论数量
			func() {
				num, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &comment.GetCommentCountReq{PostId: commonResp.Id})
				if err != nil {
					return
				}
				resp.CommentCount = num.Num
			},
			// 添加收藏数量
			func() {
				num, err := l.svcCtx.CollectRpc.GetNum(l.ctx, &collect.GetNumReq{PostId: commonResp.Id})
				if err != nil {
					return
				}
				resp.CollectionCount = num.Num
			},
		}
		// 登录用户还需要注入点赞、收藏状态
		if uid != "" {
			// 登录用户
			fns = append(fns, func() {
				// 注入收藏状态
				status, err := l.svcCtx.CollectRpc.GetStatus(l.ctx, &collect.GetStatusReq{
					PostId: commonResp.Id,
					UserId: uid,
				})
				if err != nil {
					return
				}
				resp.CollectStatus = status.Status
			}, func() {
				// 注入点赞状态
				status, err := l.svcCtx.LikeRpc.GetStatus(l.ctx, &like.GetStatusReq{PostId: commonResp.Id, UserId: uid})
				if err != nil {
					return
				}
				resp.UpvoteStatus = status.Status
			})
		}
		mr.FinishVoid(fns...)
		result = append(result, &resp)
	})

	return &types.QueryDynamicResp{List: result}, nil
}
