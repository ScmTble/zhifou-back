package like

import (
	"context"
	"git.154896.xyz/zhifou/like/like"
	"git.154896.xyz/zhifou/pkg/tool"

	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUpvoteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertUpvoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUpvoteLogic {
	return &InsertUpvoteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InsertUpvoteLogic) InsertUpvote(req *types.InsertUpvotedReq) (resp *types.InsertUpvotedResp, err error) {
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 获取当前点赞状态
	status, err := l.svcCtx.LikeRpc.GetStatus(l.ctx, &like.GetStatusReq{
		PostId: req.PostId,
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	if status.Status {
		// 为true,代表已经点赞了,则取消收藏
		_, err := l.svcCtx.LikeRpc.InsertUpvote(l.ctx, &like.InsertUpvoteReq{
			PostId: req.PostId,
			UserId: uid,
			Action: false,
		})
		if err != nil {
			return nil, err
		}
		return &types.InsertUpvotedResp{Status: false}, nil
	} else {
		// 为false,代表未点赞了,则点赞
		_, err := l.svcCtx.LikeRpc.InsertUpvote(l.ctx, &like.InsertUpvoteReq{
			PostId: req.PostId,
			UserId: uid,
			Action: true,
		})
		if err != nil {
			return nil, err
		}
		return &types.InsertUpvotedResp{Status: true}, nil
	}
}
