package logic

import (
	"context"
	"git.154896.xyz/zhifou/tag/internal/svc"
	"git.154896.xyz/zhifou/tag/model"
	"git.154896.xyz/zhifou/tag/pb"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOneLogic {
	return &FindOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// FindOne 查找一个Tag
func (l *FindOneLogic) FindOne(in *pb.FindOneReq) (*pb.FindOneResp, error) {
	var tag model.Tag

	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"_id": in.Id,
	}).One(&tag); err != nil {
		return nil, err
	}

	var resp pb.FindOneResp
	_ = copier.Copy(&resp, &tag)

	return &resp, nil
}
