package svc

import (
	"book/service/user/api/internal/config"
	"book/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 连接数据
	conn := sqlx.NewMysql(c.Mysql.DataSource)

	// 创建全局服务上下文
	// 所有的和数据打交道的可以在这里增加配置
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
