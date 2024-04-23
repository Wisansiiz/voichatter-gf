package user

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) UserInfoUpd(ctx context.Context, req *v1.UserInfoUpdReq) (res *v1.UserInfoUpdRes, err error) {
	upd, err := service.User().UserInfoUpd(ctx, model.UserInfoUpdInput{
		Username: req.Username,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UserInfoUpdRes{
		UserInfo: upd,
	}, nil
}
