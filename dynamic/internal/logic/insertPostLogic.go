package logic

import (
	"context"
	"git.154896.xyz/zhifou/idgen/idgen"
	"github.com/jinzhu/copier"
	"time"

	"git.154896.xyz/zhifou/dynamic/internal/svc"
	"git.154896.xyz/zhifou/dynamic/model"
	"git.154896.xyz/zhifou/dynamic/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type InsertPostLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInsertPostLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InsertPostLogic {
	return &InsertPostLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InsertPostLogic) InsertPost(in *pb.InsertPostReq) (*pb.InsertPostResp, error) {
	// 获取生成id
	id, err := l.svcCtx.IgRpc.IdGenServer(l.ctx, &idgen.Empty{})
	if err != nil {
		return nil, err
	}
	// 构建插入参数
	var post model.Post
	_ = copier.Copy(&post, in)
	post.CreatedTime = time.Now().Unix()
	post.Id = id.Id

	// 插入Dynamic
	if _, err := l.svcCtx.MongoDB.InsertOne(l.ctx, &post); err != nil {
		return nil, err
	}

	// 构建返回体
	var resp pb.InsertPostResp
	_ = copier.Copy(&resp, &post)

	return &resp, nil
}
