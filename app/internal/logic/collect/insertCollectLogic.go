package collect

import (
	"context"
	"git.154896.xyz/zhifou/app/internal/svc"
	"git.154896.xyz/zhifou/app/internal/types"
	"git.154896.xyz/zhifou/collect/collect"
	"git.154896.xyz/zhifou/pkg/tool"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCollectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInsertCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCollectLogic {
	return &InsertCollectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// InsertCollect 添加收藏或取消收藏
func (l *InsertCollectLogic) InsertCollect(req *types.InsertCollectReq) (*types.InsertCollectResp, error) {
	uid, err := tool.GetIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	// 获取当前状态
	status, err := l.svcCtx.CollectRpc.GetStatus(l.ctx, &collect.GetStatusReq{
		PostId: req.PostId,
		UserId: uid,
	})
	if err != nil {
		return nil, err
	}
	if status.Status {
		// 为true,代表已经收藏了,则取消收藏
		_, err := l.svcCtx.CollectRpc.InsertCollect(l.ctx, &collect.InsertCollectReq{
			PostId: req.PostId,
			UserId: uid,
			Action: false,
		})
		if err != nil {
			return nil, err
		}
		return &types.InsertCollectResp{Status: false}, nil
	} else {
		// 为false,代表未收藏了,则收藏
		_, err := l.svcCtx.CollectRpc.InsertCollect(l.ctx, &collect.InsertCollectReq{
			PostId: req.PostId,
			UserId: uid,
			Action: true,
		})
		if err != nil {
			return nil, err
		}
		return &types.InsertCollectResp{Status: true}, nil
	}
}
