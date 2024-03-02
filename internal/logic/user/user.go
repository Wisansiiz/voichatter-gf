package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
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
	err = dao.User.Ctx(ctx).
		Where(do.User{
			Username: in.Username,
		}).
		Scan(&user)
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
	_, err = dao.User.Ctx(ctx).
		Where(do.User{
			UserId: user.UserId,
		}).
		Update(do.User{
			LastLoginDate: gtime.Now(),
		})
	if err != nil {
		return token, gerror.New(`更新用户登录时间失败`)
	}
	if err = service.Session().SetUser(ctx, user); err != nil {
		return token, gerror.New(`设置用户会话失败`)
	}
	service.BizCtx().SetUser(ctx, &model.ContextUser{
		UserId:   user.UserId,
		Username: user.Username,
	})
	return
}

func (s sUser) SignOut(ctx context.Context, req *v1.SignOutReq) (res *v1.SignOutRes, err error) {
	panic("implement me")
}

func (s sUser) ServerList(ctx context.Context, req *v1.ServerListReq) (res *v1.ServerListRes, err error) {
	user := service.BizCtx().Get(ctx).User
	var servers []entity.Server
	err = g.Model("server s").
		Fields("s.*").
		LeftJoin("member m", "s.server_id = m.server_id").
		Where("m.user_id = ?", user.UserId).
		Scan(&servers)
	return &v1.ServerListRes{
		ServerList: &servers,
	}, err
}
