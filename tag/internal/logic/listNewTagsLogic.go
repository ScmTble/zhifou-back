package logic

import (
	"context"
	"git.154896.xyz/zhifou/tag/internal/svc"
	"git.154896.xyz/zhifou/tag/model"
	"git.154896.xyz/zhifou/tag/pb"
	"git.154896.xyz/zhifou/tag/tag"
	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListNewTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListNewTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListNewTagsLogic {
	return &ListNewTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListNewTagsLogic) ListNewTags(in *pb.ListNewTagReq) (*pb.ListNewTagResp, error) {
	var tags []model.Tag
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{}).Sort("-createdtime").Limit(in.Num).All(&tags); err != nil {
		return nil, err
	}

	var tagInfos []*tag.TagInfo
	_ = copier.Copy(&tagInfos, tags)

	return &pb.ListNewTagResp{
		Topics: tagInfos,
	}, nil
}
