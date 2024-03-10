// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"voichatter/api/user/v1"
)

type IUserV1 interface {
	UserId(ctx context.Context, req *v1.UserIdReq) (res *v1.UserIdRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
}
