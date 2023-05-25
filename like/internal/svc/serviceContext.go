package svc

import (
	"git.154896.xyz/zhifou/like/internal/config"
	"github.com/redis/rueidis"
)

type ServiceContext struct {
	Config config.Config
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
		Redis:  client,
	}
}
