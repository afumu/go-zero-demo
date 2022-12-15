package student

import (
	"context"

	"go-zero-demo/service/student/api/internal/svc"
	"go-zero-demo/service/student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EditLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewEditLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EditLogic {
	return &EditLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EditLogic) Edit(req *types.Student) (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	return
}
