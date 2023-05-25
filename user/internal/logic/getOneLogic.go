package logic

import (
	"context"
	"git.154896.xyz/zhifou/user/internal/svc"
	"git.154896.xyz/zhifou/user/model"
	"git.154896.xyz/zhifou/user/pb"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOneLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneLogic {
	return &GetOneLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneLogic) GetOne(in *pb.GetOneReq) (*pb.GetOneResp, error) {
	var user model.User
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"_id": in.Id,
	}).One(&user); err != nil {
		return nil, err
	}

	var resp pb.GetOneResp
	_ = copier.Copy(&resp, &user)

	return &resp, nil
}
