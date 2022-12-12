package main

import (
	"flag"
	"fmt"

	"book/service/user/api/internal/config"
	"book/service/user/api/internal/handler"
	"book/service/user/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	// 加载yml配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 读取配置文件，启动服务
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 创建全局服务上下文
	// 1、连接数据库
	// 2、
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
