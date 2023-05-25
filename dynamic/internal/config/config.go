package config

import (
	"fmt"
	"github.com/zeromicro/go-zero/zrpc"
)

// mongodb://root:123456@101.132.129.131:27017/?authSource=admin
type Config struct {
	zrpc.RpcServerConf
	MongoDB struct {
		User       string `json:",default=root,env=MONGODB_USER"`
		Password   string `json:",default=root,env=MONGODB_PASSWORD"`
		Addr       string `json:",default=192.168.211.128,env=MONGODB_ADDR"`
		Port       int    `json:",default=30017,env=MONGODB_PORT"`
		Database   string `json:",default=zhifou,env=MONGODB_DATABASE"`
		Collection string `json:",default=post,env=MONGODB_Collection"`
	}
	IgRpc zrpc.RpcClientConf
}

func (c Config) GetMongoUri() string {
	return fmt.Sprintf(
		"mongodb://%s:%s@%s:%d/?authSource=admin",
		c.MongoDB.User,
		c.MongoDB.Password,
		c.MongoDB.Addr,
		c.MongoDB.Port,
	)
}
