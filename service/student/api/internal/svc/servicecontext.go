package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/service/student/api/internal/config"
	"go-zero-demo/service/student/api/internal/middleware"
	"go-zero-demo/service/student/model"
)

type ServiceContext struct {
	Config         config.Config
	StudentModel   model.StudentModel
	TestMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:         c,
		StudentModel:   model.NewStudentModel(conn),
		TestMiddleware: middleware.NewTestMiddleware().Handle,
	}
}
