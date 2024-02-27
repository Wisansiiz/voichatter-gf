package user

import (
	"context"
	"log"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
)

type (
	sUser struct{}
)

func (s sUser) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (s sUser) SignIn(ctx context.Context, in model.UserSignInInput) (user entity.User, err error) {
	//TODO implement me
	err = dao.User.Ctx(ctx).Scan(&user)
	if err != nil {
		return entity.User{}, err
	}
	log.Println(user)
	return user, err
}

func (s sUser) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	//TODO implement me
	panic("implement me")
}

func init() {
	service.RegisterUser(New())
}
func New() service.IUser {
	return &sUser{}
}
