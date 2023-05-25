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

type ListHotTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListHotTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListHotTagsLogic {
	return &ListHotTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListHotTagsLogic) ListHotTags(in *pb.ListHotTagReq) (*pb.ListHotTagResp, error) {
	var tags []model.Tag
	if err := l.svcCtx.MongoDB.Find(l.ctx, bson.M{}).Sort("-quotenum").Limit(in.Num).All(&tags); err != nil {
		return nil, err
	}

	var tagInfos []*tag.TagInfo
	_ = copier.Copy(&tagInfos, tags)

	return &pb.ListHotTagResp{
		Topics: tagInfos,
	}, nil
}
