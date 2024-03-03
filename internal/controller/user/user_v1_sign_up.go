package user

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	_, err = service.User().SignUp(ctx, model.UserCreateInput{
		Username:          req.Username,
		Email:             req.Email,
		PasswordHash:      req.PasswordHash,
		ReenteredPassword: req.ReenteredPassword,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SignUpRes{}, nil
}
