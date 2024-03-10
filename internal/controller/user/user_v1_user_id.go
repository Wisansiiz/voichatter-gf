package user

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserId(ctx context.Context, req *v1.UserIdReq) (res *v1.UserIdRes, err error) {
	userId, err := service.User().UserId(ctx, req)
	if err != nil {
		return nil, err
	}
	return userId, nil
}
