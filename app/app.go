package main

import (
	"flag"
	"fmt"
	"git.154896.xyz/zhifou/app/internal/config"
	"git.154896.xyz/zhifou/app/internal/handler"
	"git.154896.xyz/zhifou/app/internal/svc"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/app.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())

	server := rest.MustNewServer(
		c.RestConf,
		rest.WithCors(),
	)
	defer server.Stop()

	//httpx.SetValidator(validatorx.NewCusValidator())

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.PrintRoutes()
	server.Start()
}
