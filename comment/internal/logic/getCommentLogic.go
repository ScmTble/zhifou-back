package logic

import (
	"context"
	"git.154896.xyz/zhifou/comment/internal/svc"
	"git.154896.xyz/zhifou/comment/model"
	"git.154896.xyz/zhifou/comment/pb"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCommentLogic {
	return &GetCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetCommentLogic) GetComment(in *pb.GetCommentReq) (*pb.GetCommentResp, error) {
	var comments []model.Comment
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{
		"postid":      in.PostId,
		"deletedtime": 0,
	}).Sort("-createdtime").All(&comments); err != nil {
		return nil, err
	}

	var commentsResp []*pb.CommonCommentResp
	_ = copier.Copy(&commentsResp, &commentsResp)

	return &pb.GetCommentResp{Comments: commentsResp}, nil
}
