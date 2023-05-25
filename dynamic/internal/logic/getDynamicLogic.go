package logic

import (
	"context"
	"git.154896.xyz/zhifou/dynamic/model"
	"github.com/jinzhu/copier"
	"github.com/qiniu/qmgo/operator"
	"github.com/zeromicro/go-zero/core/fx"
	"go.mongodb.org/mongo-driver/bson"

	"git.154896.xyz/zhifou/dynamic/internal/svc"
	"git.154896.xyz/zhifou/dynamic/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDynamicLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDynamicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDynamicLogic {
	return &GetDynamicLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetDynamic 分页获取动态
func (l *GetDynamicLogic) GetDynamic(in *pb.GetDynamicReq) (*pb.GetDynamicResp, error) {
	var posts []model.Post
	// 分页构建查询条件
	query := bson.M{
		"_id": bson.M{
			operator.Lt: in.LastId,
		},
	}

	if err := l.svcCtx.MongoDB.
		Find(l.ctx, query).
		Sort("-_id").
		Limit(in.PageNum).
		All(&posts); err != nil {
		return nil, err
	}

	// 转换数据
	var result []*pb.CommonResp
	fx.From(func(source chan<- any) {
		for _, v := range posts {
			source <- v
		}
	}).ForEach(func(item any) {
		post := item.(model.Post)
		var resp pb.CommonResp
		_ = copier.Copy(&resp, &post)
		result = append(result, &resp)
	})

	return &pb.GetDynamicResp{
		Data: result,
	}, nil
}
