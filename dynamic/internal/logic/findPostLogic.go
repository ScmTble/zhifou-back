package logic

import (
	"context"
	"git.154896.xyz/zhifou/dynamic/internal/svc"
	"git.154896.xyz/zhifou/dynamic/model"
	"git.154896.xyz/zhifou/dynamic/pb"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPostLogic {
	return &FindPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindPostLogic) FindPost(in *pb.FindOneReq) (*pb.FindOneResp, error) {
	var post model.Post
	// 查询数据
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.D{
		{
			"_id", in.Id,
		},
	}).One(&post); err != nil {
		return nil, err
	}

	// 转换数据
	var resp pb.FindOneResp
	_ = copier.Copy(&resp, &post)

	return &resp, nil
}
