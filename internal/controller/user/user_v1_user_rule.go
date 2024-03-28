package user

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserRole(ctx context.Context, req *v1.UserRoleReq) (res *v1.UserRoleRes, err error) {
	user, err := service.User().UserRole(ctx, model.ModifyUserRoleInput{
		UserId:       req.UserId,
		ServerId:     req.ServerId,
		SPermissions: req.SPermissions,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
