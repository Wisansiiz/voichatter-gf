package user

import (
	"context"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	token, err := service.User().SignIn(ctx, model.UserSignInInput{
		Username:     req.Username,
		PasswordHash: req.PasswordHash,
	})
	if err != nil {
		return nil, err
	}
	return &v1.SignInRes{
		Token: token,
	}, nil
}
