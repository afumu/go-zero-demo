package student

import (
	"context"
	"fmt"
	"go-zero-demo/service/student/model"

	"go-zero-demo/service/student/api/internal/svc"
	"go-zero-demo/service/student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddLogic {
	return &AddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddLogic) Add(req *types.Student) (resp *types.Result, err error) {

	logx.Info("add student.....")
	// todo: add your logic here and delete this line
	m := &model.Student{Name: req.Name}
	result, err := l.svcCtx.StudentModel.Insert(l.ctx, m)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	message := fmt.Sprintf("成功了:%v", id)
	return &types.Result{Data: message}, nil
}
