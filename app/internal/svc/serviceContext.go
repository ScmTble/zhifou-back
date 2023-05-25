package svc

import (
	"git.154896.xyz/zhifou/app/internal/config"
	"git.154896.xyz/zhifou/app/internal/middleware"
	"git.154896.xyz/zhifou/collect/collect"
	"git.154896.xyz/zhifou/comment/comment"
	"git.154896.xyz/zhifou/dynamic/dynamic"
	"git.154896.xyz/zhifou/gorse/gorse"
	"git.154896.xyz/zhifou/like/like"
	"git.154896.xyz/zhifou/tag/tag"
	"git.154896.xyz/zhifou/user/user"

	"github.com/nats-io/nats.go"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	JwtAuth    rest.Middleware
	UserRpc    user.User
	DynamicRpc dynamic.Dynamic
	TagRpc     tag.Tag
	CommentRpc comment.Comment
	GorseRpc   gorse.Gorse
	CollectRpc collect.Collect
	LikeRpc    like.Like
	Nc         *nats.Conn
}

func NewServiceContext(c config.Config) *ServiceContext {
	nc, err := nats.Connect(c.Nats.Url)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		JwtAuth: middleware.NewJwtAuthMiddleware(c.JwtAuth.AccessSecret).Handle,

		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpcConf)),
		DynamicRpc: dynamic.NewDynamic(zrpc.MustNewClient(c.DynamicRpcConf)),
		TagRpc:     tag.NewTag(zrpc.MustNewClient(c.TagRpcConf)),
		CommentRpc: comment.NewComment(zrpc.MustNewClient(c.CommentRpcConf)),
		CollectRpc: collect.NewCollect(zrpc.MustNewClient(c.CollectRpcConf)),
		LikeRpc:    like.NewLike(zrpc.MustNewClient(c.LikeRpcConf)),
		GorseRpc:   gorse.NewGorse(zrpc.MustNewClient(c.GorseRpcConf)),
		Nc:         nc,
	}
}
