package logic

import (
	"context"
	"fmt"
	"git.154896.xyz/zhifou/collect/internal/svc"
	"git.154896.xyz/zhifou/collect/pb"
	"git.154896.xyz/zhifou/pkg/consts"
	"github.com/redis/rueidis"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStatusLogic {
	return &GetStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetStatus 用户是否收藏了该动态
func (l *GetStatusLogic) GetStatus(in *pb.GetStatusReq) (*pb.GetStatusResp, error) {
	userKey := fmt.Sprintf(consts.CollectPrefix+"%s", in.UserId)
	client := l.svcCtx.Redis
	// 查询Redis
	// 判断key是否存在
	result := client.Do(
		l.ctx,
		client.B().Zscore().Key(userKey).Member(in.PostId).Build(),
	)

	switch result.Error() {
	case nil:
		return &pb.GetStatusResp{Status: true}, nil
	case rueidis.Nil:
		// redis返回的消息为空
		return &pb.GetStatusResp{Status: false}, nil
	default:
		return nil, fmt.Errorf("redis result is %v", result)
	}
}
