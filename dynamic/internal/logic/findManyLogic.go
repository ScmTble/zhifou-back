package logic

import (
	"context"
	"git.154896.xyz/zhifou/dynamic/internal/svc"
	"git.154896.xyz/zhifou/dynamic/model"
	"git.154896.xyz/zhifou/dynamic/pb"
	"github.com/jinzhu/copier"
	"github.com/qiniu/qmgo/operator"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"go.mongodb.org/mongo-driver/bson"
)

type FindManyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindManyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindManyLogic {
	return &FindManyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindManyLogic) FindMany(in *pb.FindManyReq) (*pb.FindManyReqResp, error) {
	var posts []model.Post
	// 构建查询条件
	query := bson.D{
		{
			"_id", bson.D{
				{
					operator.In, in.Ids,
				},
			},
		},
	}
	// 获取数据
	if err := l.svcCtx.MongoDB.Find(l.ctx, query).All(&posts); err != nil {
		return nil, err
	}

	// 转换数据
	result, err := mr.MapReduce(func(source chan<- model.Post) {
		for _, v := range posts {
			source <- v
		}
	}, func(item model.Post, writer mr.Writer[*pb.CommonResp], cancel func(error)) {
		var resp pb.CommonResp
		_ = copier.Copy(&resp, &item)
		writer.Write(&resp)
	}, func(pipe <-chan *pb.CommonResp, writer mr.Writer[[]*pb.CommonResp], cancel func(error)) {
		var result []*pb.CommonResp
		for v := range pipe {
			result = append(result, v)
		}
		writer.Write(result)
	})
	if err != nil {
		return nil, err
	}

	return &pb.FindManyReqResp{
		Data: result,
	}, nil
}
