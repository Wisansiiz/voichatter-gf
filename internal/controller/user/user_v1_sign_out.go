package user

import (
	"context"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) SignOut(ctx context.Context, req *v1.SignOutReq) (err error) {
	err = service.User().SignOut(ctx, req)
	return
}
