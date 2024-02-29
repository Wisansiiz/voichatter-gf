package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/dao"
	"voichatter/internal/logic/jwt"
	"voichatter/internal/model"
	"voichatter/internal/model/do"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}
func New() service.IUser {
	return &sUser{}
}

func (s sUser) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	//TODO implement me
	panic("implement me")
}

func (s sUser) SignIn(ctx context.Context, in model.UserSignInInput) (token string, err error) {
	var user *entity.User
	err = dao.User.Ctx(ctx).Where(do.User{
		Username: in.Username,
	}).Scan(&user)
	if err != nil {
		return token, gerror.New(`账号或密码错误`)
	}
	if user == nil {
		return token, gerror.New(`账号或密码错误`)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(in.PasswordHash)); err != nil {
		return token, gerror.New(`账号或密码错误`)
	}
	token, err = jwt.GenerateToken(*user)
	if err != nil {
		return token, gerror.New(`账号或密码错误`)
	}
	_, err = dao.User.Ctx(ctx).Where(do.User{
		UserId: user.UserId,
	}).Update(do.User{
		LastLoginDate: gtime.Now(),
	})
	if err != nil {
		return token, gerror.New(`更新用户登录时间失败`)
	}
	return
}

func (s sUser) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	panic("implement me")
}
