package logic

import (
	"context"

	"git.154896.xyz/zhifou/tag/internal/svc"
	"git.154896.xyz/zhifou/tag/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListTagsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListTagsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListTagsLogic {
	return &ListTagsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListTagsLogic) ListTags(in *pb.ListTagReq) (*pb.ListTagResp, error) {
	// todo: add your logic here and delete this line

	return &pb.ListTagResp{}, nil
}
