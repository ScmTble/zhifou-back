package logic

import (
	"context"
	"git.154896.xyz/zhifou/like/internal/svc"
	"git.154896.xyz/zhifou/like/pb"
	"git.154896.xyz/zhifou/pkg/consts"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertUpvoteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertUpvoteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertUpvoteLogic {
	return &InsertUpvoteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// InsertUpvote 点赞或取消点赞
func (l *InsertUpvoteLogic) InsertUpvote(in *pb.InsertUpvoteReq) (*pb.InsertUpvoteResp, error) {
	client := l.svcCtx.Redis
	userKey := consts.UpvotePrefix + in.UserId
	postKey := consts.PostPrefix + in.PostId
	timestamp := float64(time.Now().Unix())
	// 点赞
	if in.Action {
		// ZSet 添加
		if err := client.Do(
			l.ctx,
			client.B().Zadd().Key(userKey).Ch().ScoreMember().ScoreMember(timestamp, in.PostId).Build(),
		).Error(); err != nil {
			return nil, err
		}
		// 动态点赞数加1
		if err := client.Do(l.ctx,
			client.B().Hincrby().Key(postKey).Field(consts.PostUpvoteCountKey).Increment(1).Build(),
		).Error(); err != nil {
			return nil, err
		}
		return &pb.InsertUpvoteResp{Success: true}, nil
	} else {
		// 取消点赞
		// ZSet 删除
		if err := client.Do(
			l.ctx,
			client.B().Zrem().Key(userKey).Member(in.PostId).Build(),
		).Error(); err != nil {
			return nil, err
		}
		// 动态点赞数减1
		if err := client.Do(l.ctx,
			client.B().Hincrby().Key(postKey).Field(consts.PostUpvoteCountKey).Increment(-1).Build(),
		).Error(); err != nil {
			return nil, err
		}
		return &pb.InsertUpvoteResp{Success: true}, nil
	}
}
