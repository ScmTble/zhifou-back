package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Rueidis struct {
		Address []string
		DB      int
	}
}
