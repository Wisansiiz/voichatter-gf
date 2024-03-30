package user

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserAvatar(ctx context.Context, req *v1.UserAvatarReq) (res *v1.UserAvatarRes, err error) {
	avatar, err := service.User().UserAvatar(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return avatar, nil
}
