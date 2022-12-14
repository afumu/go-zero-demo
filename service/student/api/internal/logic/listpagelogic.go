package logic

import (
	"context"

	"book/service/student/api/internal/svc"
	"book/service/student/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPageLogic {
	return &ListPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListPageLogic) ListPage() (resp *types.LoginRep, err error) {
	// todo: add your logic here and delete this line

	return
}
