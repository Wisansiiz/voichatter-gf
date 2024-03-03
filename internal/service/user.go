// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/model"
)

type (
	IUser interface {
		SignUp(ctx context.Context, in model.UserCreateInput) (res *v1.SignUpRes, err error)
		SignIn(ctx context.Context, in model.UserSignInInput) (token string, err error)
		SignOut(ctx context.Context, _ *v1.SignOutReq) (err error)
		ServerList(ctx context.Context, _ *v1.ServerListReq) (res *v1.ServerListRes, err error)
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
