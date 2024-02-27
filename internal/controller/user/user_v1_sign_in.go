package user

import (
	"context"
	"log"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"voichatter/api/user/v1"
)

func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	in, err := service.User().SignIn(ctx, model.UserSignInInput{
		Password: req.Password,
		Passport: req.Passport,
	})
	if err != nil {
		return nil, err
	}
	log.Println("signIn", req)
	res = &v1.SignInRes{
		User: &in,
	}
	return
}
