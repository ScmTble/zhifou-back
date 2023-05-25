package logic

import (
	"context"
	"git.154896.xyz/zhifou/idgen/idgen"
	"git.154896.xyz/zhifou/tag/internal/svc"
	"git.154896.xyz/zhifou/tag/model"
	"git.154896.xyz/zhifou/tag/pb"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertTagLogic {
	return &InsertTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertTagLogic) InsertTag(in *pb.InsertTagReq) (*pb.InsertTagResp, error) {
	result, err := l.svcCtx.IgRpc.IdGenServer(l.ctx, &idgen.Empty{})
	if err != nil {
		return nil, err
	}

	tag := model.Tag{
		Id:          result.Id,
		Tag:         in.Tag,
		Avatar:      in.Avatar,
		QuoteNum:    0,
		CreatedTime: time.Now().Unix(),
		UpdatedTime: 0,
		DeletedTime: 0,
	}

	_, err = l.svcCtx.MongoDB.InsertOne(l.ctx, &tag)
	if err != nil {
		return nil, err
	}

	return &pb.InsertTagResp{
		Id: tag.Id,
	}, nil
}
