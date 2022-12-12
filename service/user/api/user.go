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

	// 一、加载yml配置文件
	var c config.Config
	conf.MustLoad(*configFile, &c)

	// 二、根据配置文件，创建server
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 三、创建全局服务上下文
	// 	1、连接数据库
	// 	2、创建操作数据库的Model
	ctx := svc.NewServiceContext(c)

	// 四、注册处理器
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	// 五、启动服务
	server.Start()
}
