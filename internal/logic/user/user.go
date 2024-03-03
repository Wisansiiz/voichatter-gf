package user

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"golang.org/x/crypto/bcrypt"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
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

func (s *sUser) SignUp(ctx context.Context, in model.UserCreateInput) (res *v1.SignUpRes, err error) {
	pwd, err := bcrypt.GenerateFromPassword([]byte(in.PasswordHash), bcrypt.DefaultCost) //加密处理
	if err != nil {
		return res, gerror.New(`密码加密失败`)
	}
	var user = entity.User{
		Username:         in.Username,
		Email:            in.Email,
		PasswordHash:     string(pwd),
		RegistrationDate: gtime.Now(),
	}
	if _, err = dao.User.Ctx(ctx).Insert(user); err != nil {
		return res, gerror.New(`注册失败`)
	}
	return
}
func (s *sUser) UserList(ctx context.Context, serverId uint64) (res *v1.UserListRes, err error) {
	// 获取服务器的成员列表的用户ID
	userIds, err := dao.Member.Ctx(ctx).
		Fields("user_id").
		Where("server_id = ?", serverId).
		Array()
	if err != nil {
		return nil, gerror.New("获取用户列表失败")
	}
	// 若找不到成员则直接返回空列表
	if len(userIds) == 0 {
		return nil, nil
	}
	var users []model.UserList4Server
	// 使用查询到的用户ID获取用户信息列表
	err = g.Model("user").
		Fields("user.user_id", "user.username", "user.email", "user.avatar_url", "member.s_permissions", "user.last_login_date").
		LeftJoin("member", "user.user_id = member.user_id").
		Where("user.user_id IN(?) AND member.server_id = ?", userIds, serverId).
		Scan(&users)
	if err != nil {
		return nil, gerror.New("获取用户信息列表失败")
	}
	return &v1.UserListRes{
		Users: &users,
	}, nil
}
