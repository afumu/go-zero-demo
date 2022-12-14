package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"

	"book/service/user/api/internal/svc"
	"book/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAllUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindAllUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAllUserLogic {
	return &FindAllUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 返回数据库的所有数据
func (l *FindAllUserLogic) FindAllUser() (resp *types.Result, err error) {
	// todo: add your logic here and delete this line

	usersFromDb, err := l.svcCtx.UserModel.FindAll(l.ctx)
	if err != nil {
		return nil, err
	}
	var users []*types.User

	for _, v := range usersFromDb {
		user := &types.User{}
		err = copier.Copy(&user, &v)
		user.CreateTime = v.CreateTime.UnixMilli()
		user.UpdateTime = v.UpdateTime.UnixMilli()
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, user)
	}

	return &types.Result{User: users}, nil

}
