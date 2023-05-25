package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	IgRpc   zrpc.RpcClientConf
	Rueidis struct {
		Address []string
		DB      int
	}
}
