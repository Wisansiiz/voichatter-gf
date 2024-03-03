package user

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserList(ctx context.Context, req *v1.UserListReq) (res *v1.UserListRes, err error) {
	list, err := service.User().UserList(ctx, req.ServerId)
	if err != nil {
		return nil, err
	}
	return list, nil
}
