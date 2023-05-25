package logic

import (
	"context"
	"git.154896.xyz/zhifou/comment/internal/svc"
	"git.154896.xyz/zhifou/comment/model"
	"git.154896.xyz/zhifou/comment/pb"
	"git.154896.xyz/zhifou/idgen/idgen"
	"github.com/jinzhu/copier"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

// NewInsertCommentLogic 新增评论
func NewInsertCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertCommentLogic {
	return &InsertCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertCommentLogic) InsertComment(in *pb.InsertCommentReq) (*pb.InsertCommentResp, error) {
	resutl, err := l.svcCtx.IgRpc.IdGenServer(l.ctx, &idgen.Empty{})
	if err != nil {
		return nil, err
	}

	comment := model.Comment{
		Id:          resutl.Id,
		PostId:      in.PostId,
		UserId:      in.UserId,
		Content:     in.Content,
		CreatedTime: time.Now().Unix(),
		UpdatedTime: 0,
		DeletedTime: 0,
	}

	_, err = l.svcCtx.MongoDB.InsertOne(l.ctx, &comment)
	if err != nil {
		return nil, err
	}

	var resp pb.InsertCommentResp

	_ = copier.Copy(&resp, &comment)

	return &resp, nil
}
