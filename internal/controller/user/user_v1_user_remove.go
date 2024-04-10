package user

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserRemove(ctx context.Context, req *v1.UserRemoveReq) (res *v1.UserRemoveRes, err error) {
	err = service.User().UserRemove(ctx, model.UserRemoveInput{
		ServerId: req.ServerId,
		UserId:   req.UserId,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
