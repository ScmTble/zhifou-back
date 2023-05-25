package logic

import (
	"context"
	"fmt"
	"git.154896.xyz/zhifou/collect/internal/svc"
	"git.154896.xyz/zhifou/collect/pb"
	"git.154896.xyz/zhifou/pkg/consts"
	"strconv"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCollectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCollectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCollectLogic {
	return &GetCollectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetCollect 分页获取所有的收藏
func (l *GetCollectLogic) GetCollect(in *pb.GetCollectReq) (*pb.GetCollectResp, error) {
	client := l.svcCtx.Redis
	userKey := fmt.Sprintf(consts.CollectPrefix+"%s", in.UserId)
	max := strconv.FormatInt(time.Now().Unix(), 10)
	result := client.Do(l.ctx,
		client.B().Zrevrangebyscore().Key(userKey).Max(max).Min("0").Limit(in.Offset, in.Num).Build(),
	)
	if result.Error() != nil {
		return nil, result.Error()
	}

	ids, err := result.AsStrSlice()
	if err != nil {
		return nil, err
	}

	return &pb.GetCollectResp{Ids: ids}, nil
}
