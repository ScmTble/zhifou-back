package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	UserRpcConf    zrpc.RpcClientConf
	DynamicRpcConf zrpc.RpcClientConf
	TagRpcConf     zrpc.RpcClientConf
	CommentRpcConf zrpc.RpcClientConf
	GorseRpcConf   zrpc.RpcClientConf
	CollectRpcConf zrpc.RpcClientConf
	LikeRpcConf    zrpc.RpcClientConf
	JwtAuth        struct {
		AccessSecret string
		AccessExpire int64
	}
	Nats struct {
		Url string
	}
}
