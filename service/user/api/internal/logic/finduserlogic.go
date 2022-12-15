package logic

import (
	"context"
	"fmt"

	"go-zero-demo/service/user/api/internal/svc"
	"go-zero-demo/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser(req *types.UserParam) (*types.User, error) {
	// todo: add your logic here and delete this line

	user, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Username)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	t := &types.User{
		Name:     user.Name,
		Password: user.Password,
	}
	return t, nil
}
