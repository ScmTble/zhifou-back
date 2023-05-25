package svc

import (
	"git.154896.xyz/zhifou/collect/internal/config"
	"git.154896.xyz/zhifou/idgen/idgen"
	"github.com/redis/rueidis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config
	IgRpc  idgen.IdGen
	Redis  rueidis.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	client, err := rueidis.NewClient(rueidis.ClientOption{
		InitAddress: c.Rueidis.Address,
		SelectDB:    c.Rueidis.DB,
	})
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		IgRpc:  idgen.NewIdGen(zrpc.MustNewClient(c.IgRpc)),
		Redis:  client,
	}
}
