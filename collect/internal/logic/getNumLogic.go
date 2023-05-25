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

type GetNumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNumLogic {
	return &GetNumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetNum 获取某条动态的收藏数量
func (l *GetNumLogic) GetNum(in *pb.GetNumReq) (*pb.GetNumResp, error) {
	postKey := fmt.Sprintf(consts.PostPrefix+"%s", in.PostId)
	client := l.svcCtx.Redis
	result := client.Do(
		l.ctx,
		client.B().Hget().Key(postKey).Field(consts.PostCollectCountKey).Build(),
	)
	switch result.Error() {
	case nil:
		v, err := result.AsInt64()
		if err != nil {
			return nil, err
		}
		return &pb.GetNumResp{
			Num: v,
		}, nil
	case rueidis.Nil:
		// redis返回的消息为空
		return &pb.GetNumResp{
			Num: 0,
		}, nil
	default:
		return nil, fmt.Errorf("redis result is %v", result)
	}
}
