package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"iotPlatform/user/api/internal/config"
	"iotPlatform/user/api/internal/handler"
	"iotPlatform/user/api/internal/svc"
)

// var configFile = flag.String("f", "user/api/etc/user-api.yaml", "the config file")
var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
