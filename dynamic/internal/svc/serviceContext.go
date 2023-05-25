package svc

import (
	"context"
	"git.154896.xyz/zhifou/dynamic/internal/config"
	"git.154896.xyz/zhifou/idgen/idgen"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/qiniu/qmgo"
)

type ServiceContext struct {
	Config  config.Config
	MongoDB *qmgo.Collection
	IgRpc   idgen.IdGen
}

func NewServiceContext(c config.Config) *ServiceContext {
	ctx := context.Background()
	cli, err := qmgo.NewClient(ctx, &qmgo.Config{
		Uri: c.GetMongoUri(),
	})
	if err != nil {
		panic(err)
	}

	collection := cli.Database(c.MongoDB.Database).Collection(c.MongoDB.Collection)
	return &ServiceContext{
		Config:  c,
		MongoDB: collection,
		IgRpc:   idgen.NewIdGen(zrpc.MustNewClient(c.IgRpc)),
	}
}