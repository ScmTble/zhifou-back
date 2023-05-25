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

type GetOneByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOneByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOneByNameLogic {
	return &GetOneByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOneByNameLogic) GetOneByName(in *pb.GetOneByNameReq) (*pb.GetOneByNameResp, error) {
	var user model.User
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"username": in.UserName,
	}).One(&user); err != nil {
		return nil, err
	}

	var resp pb.GetOneByNameResp
	_ = copier.Copy(&resp, &user)

	return &resp, nil
}
