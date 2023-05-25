package logic

import (
	"context"
	"fmt"
	"git.154896.xyz/zhifou/collect/internal/svc"
	"git.154896.xyz/zhifou/collect/pb"
	"git.154896.xyz/zhifou/pkg/consts"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type InsertCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCollectLogic {
	return &InsertCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsertCollect 新增收藏
func (l *InsertCollectLogic) InsertCollect(in *pb.InsertCollectReq) (*pb.InsertCollectResp, error) {
	client := l.svcCtx.Redis
	userKey := fmt.Sprintf(consts.CollectPrefix+"%s", in.UserId)
	postKey := fmt.Sprintf(consts.PostPrefix+"%s", in.PostId)
	timestamp := float64(time.Now().Unix())
	// 添加收藏
	if in.Action {
		// ZSet 添加
		if err := client.Do(
			l.ctx,
			client.B().Zadd().Key(userKey).Ch().ScoreMember().ScoreMember(timestamp, in.PostId).Build(),
		).Error(); err != nil {
			return nil, err
		}
		// 动态收藏数加1
		if err := client.Do(l.ctx,
			client.B().Hincrby().Key(postKey).Field(consts.PostCollectCountKey).Increment(1).Build(),
		).Error(); err != nil {
			return nil, err
		}

		return &pb.InsertCollectResp{Success: true}, nil
	} else {
		// 取消收藏
		// ZSet 删除
		if err := client.Do(
			l.ctx,
			client.B().Zrem().Key(userKey).Member(in.PostId).Build(),
		).Error(); err != nil {
			return nil, err
		}
		// 动态收藏数减1
		if err := client.Do(l.ctx,
			client.B().Hincrby().Key(postKey).Field(consts.PostCollectCountKey).Increment(-1).Build(),
		).Error(); err != nil {
			return nil, err
		}

		return &pb.InsertCollectResp{Success: true}, nil
	}
}
