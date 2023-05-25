package collect

import (
	"context"
	"git.154896.xyz/zhifou/collect/collect"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/dynamic/dynamic"
	"git.154896.xyz/zhifou/like/like"
	"git.154896.xyz/zhifou/pkg/tool"
	userRpc "git.154896.xyz/zhifou/user/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectLogic {
	return &GetCollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetCollectLogic) GetCollect(req *types.GetCollectReq) (*types.GetCollectResp, error) {
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	ids, err := l.svcCtx.CollectRpc.GetCollect(l.ctx, &collect.GetCollectReq{
		UserId: uid,
		Offset: req.Offset,
		Num:    req.Num,
	})
	if err != nil {
		return nil, err
	}

	if len(ids.Ids) == 0 {
		return nil, nil
	}

	// 获取收藏的动态详细信息
	dynamics, err := l.svcCtx.DynamicRpc.FindMany(l.ctx, &dynamic.FindManyReq{
		Ids: ids.Ids,
	})
	if err != nil {
		return nil, err
	}

	result, err := mr.MapReduce(func(source chan<- *dynamic.CommonResp) {
		for _, v := range dynamics.Data {
			source <- v
		}
	}, func(item *dynamic.CommonResp, writer mr.Writer[*types.CommenPostResp], cancel func(error)) {
		var resp types.CommenPostResp
		_ = copier.Copy(&resp, item)

		fns := []func(){
			// 注入用户信息
			func() {
				var user types.CommonUserResp
				u, _ := l.svcCtx.UserRpc.GetOne(l.ctx, &userRpc.GetOneReq{Id: item.UserId})
				_ = copier.Copy(&user, u)
				resp.User = &user
			},
			//注入点赞数量
			func() {
				num, err := l.svcCtx.LikeRpc.GetNum(l.ctx, &like.GetNumReq{PostId: item.Id})
				if err != nil {
					return
				}
				resp.UpvoteCount = num.Num
			},
			// 添加评论数量
			func() {
				num, err := l.svcCtx.CommentRpc.GetCommentCount(l.ctx, &comment.GetCommentCountReq{PostId: item.Id})
				if err != nil {
					return
				}
				resp.CommentCount = num.Num
			},
			// 添加收藏数量
			func() {
				num, err := l.svcCtx.CollectRpc.GetNum(l.ctx, &collect.GetNumReq{PostId: item.Id})
				if err != nil {
					return
				}
				resp.CollectionCount = num.Num
			},
			func() {
				// 注入收藏状态
				resp.CollectStatus = true
			}, func() {
				// 注入点赞状态
				status, err := l.svcCtx.LikeRpc.GetStatus(l.ctx, &like.GetStatusReq{PostId: item.Id, UserId: uid})
				if err != nil {
					return
				}
				resp.UpvoteStatus = status.Status
			},
		}
		mr.FinishVoid(fns...)
		writer.Write(&resp)

	}, func(pipe <-chan *types.CommenPostResp, writer mr.Writer[[]*types.CommenPostResp], cancel func(error)) {
		var result []*types.CommenPostResp
		for v := range pipe {
			result = append(result, v)
		}
		writer.Write(result)
	})

	return &types.GetCollectResp{List: result}, nil
}
