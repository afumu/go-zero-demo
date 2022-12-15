package main

import (
	"flag"
	"fmt"
	middleware "go-zero-demo/service/student/common"

	"go-zero-demo/service/student/api/internal/config"
	"go-zero-demo/service/student/api/internal/handler"
	"go-zero-demo/service/student/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/student-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)

	server.Use(middleware.NewCommonJwtAuthMiddleware().Handle)

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
