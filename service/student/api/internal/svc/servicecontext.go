package svc

import (
	"book/service/student/api/internal/config"
	"book/service/student/api/internal/middleware"
	"book/service/student/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
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
