// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"voichatter/api/user/v1"
)

type IUserV1 interface {
	UserRole(ctx context.Context, req *v1.UserRoleReq) (res *v1.UserRoleRes, err error)
	UserRemove(ctx context.Context, req *v1.UserRemoveReq) (res *v1.UserRemoveRes, err error)
	UserId(ctx context.Context, req *v1.UserIdReq) (res *v1.UserIdRes, err error)
	UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error)
	SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error)
	UserAvatar(ctx context.Context, req *v1.UserAvatarReq) (res *v1.UserAvatarRes, err error)
}
