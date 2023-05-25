package logic

import (
	"context"
	"github.com/yitter/idgenerator-go/idgen"
	"strconv"

	"git.154896.xyz/zhifou/idgen/internal/svc"
	"git.154896.xyz/zhifou/idgen/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdGenServerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIdGenServerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdGenServerLogic {
	return &IdGenServerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *IdGenServerLogic) IdGenServer(in *pb.Empty) (*pb.IdGenResp, error) {
	options := idgen.NewIdGeneratorOptions(l.svcCtx.Config.Node)
	idgen.SetIdGenerator(options)
	id := idgen.NextId()
	return &pb.IdGenResp{
		Id: strconv.FormatInt(id, 10),
	}, nil
}
