package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"

	"git.154896.xyz/zhifou/comment/internal/svc"
	"git.154896.xyz/zhifou/comment/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentCountLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentCountLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentCountLogic {
	return &GetCommentCountLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentCountLogic) GetCommentCount(in *pb.GetCommentCountReq) (*pb.GetCommentCountResp, error) {
	n, err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"postid": in.PostId,
	}).Count()
	if err != nil {
		return nil, err
	}

	return &pb.GetCommentCountResp{
		Num: n,
	}, nil
}
