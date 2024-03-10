// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IUser interface {
		SignUp(ctx context.Context, in model.UserCreateInput) (res *v1.SignUpRes, err error)
		UserList(ctx context.Context, serverId uint64) (res *v1.UserListRes, err error)
		LoginFunc(r *ghttp.Request) (string, interface{})
		UserId(ctx context.Context, _ *v1.UserIdReq) (res *v1.UserIdRes, err error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
