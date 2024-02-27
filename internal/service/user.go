// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
)

type (
	IUser interface {
		SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
		SignIn(ctx context.Context, in model.UserSignInInput) (user entity.User, err error)
		SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error)
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
