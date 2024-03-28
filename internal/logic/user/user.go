package user

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
	v1 "voichatter/api/user/v1"
	"voichatter/internal/consts"
	"voichatter/internal/dao"
	"voichatter/internal/logic/cache"
	"voichatter/internal/model"
	"voichatter/internal/model/do"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
	"voichatter/utility/errResponse"
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
	usersKey := fmt.Sprintf("%s-%d", consts.ServerUsers, serverId)
	usersValue, err := g.Redis().Get(ctx, usersKey)
	if err != nil {
		return nil, errResponse.DbOperationError("获取用户信息列表失败")
	}
	var users []*model.UserList4Server
	if err = usersValue.Struct(&users); err != nil {
		return nil, errResponse.OperationFailed("获取用户信息列表失败")
	}
	if users != nil {
		return &v1.UserListRes{
			Users: users,
		}, nil
	}

	// 获取服务器的成员列表的用户ID
	userIds, err := dao.Member.Ctx(ctx).
		Fields("user_id").
		Where("server_id = ?", serverId).
		Array()
	if err != nil {
		return nil, errResponse.DbOperationError("获取用户信息列表失败")
	}

	// 使用查询到的用户ID获取用户信息列表
	err = dao.User.Ctx(ctx).
		Fields("user.user_id", "user.username", "user.email", "user.avatar_url", "member.s_permissions", "user.last_login_date").
		LeftJoin("member", "user.user_id = member.user_id").
		Where("user.user_id IN(?) AND member.server_id = ?", userIds, serverId).
		OrderAsc("member.join_date").
		Scan(&users)
	if err != nil {
		return nil, errResponse.DbOperationError("获取用户信息列表失败")
	}

	if err = g.Redis().SetEX(ctx, usersKey, users, int64(gtime.D)); err != nil {
		return nil, errResponse.DbOperationError("设置失败")
	}

	return &v1.UserListRes{
		Users: users,
	}, nil
}

func (s *sUser) LoginFunc(r *ghttp.Request) (string, interface{}) {
	var u *entity.User
	var in *model.UserSignInInput
	if err := r.Parse(&in); err != nil {
		r.SetError(err)
		r.Exit()
	}
	err := g.DB().Model(entity.User{}).
		Where(do.User{
			Username: in.Username,
		}).
		Scan(&u)
	if err != nil || u == nil {
		r.SetError(gerror.New("账号或密码错误"))
		r.Exit()
	}
	if err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(in.PasswordHash)); err != nil {
		r.SetError(gerror.New("账号或密码错误"))
		r.Exit()
	}
	_, err = g.DB().Model(entity.User{}).
		Where(do.User{
			UserId: u.UserId,
		}).
		Update(do.User{
			LastLoginDate: gtime.Now(),
		})
	// 唯一标识，扩展参数user data
	return gconv.String(u.UserId), &u
}

func (s *sUser) UserId(ctx context.Context, _ *v1.UserIdReq) (res *v1.UserIdRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	return &v1.UserIdRes{
		UserId: userId,
	}, nil
}

func (s *sUser) UserRole(ctx context.Context, in model.ModifyUserRoleInput) (res *v1.UserRoleRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Server.Ctx(ctx).Where("server_id = ? AND creator_user_id = ?", in.ServerId, userId).Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限不足")
	}
	if in.UserId == userId {
		return nil, errResponse.OperationFailed("不能修改自己的权限")
	}
	count, err = dao.Permission.Ctx(ctx).
		Where("permission_type = ?", in.SPermissions).
		Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限参数错误")
	}
	_, err = dao.Member.Ctx(ctx).
		Fields("s_permissions").
		Data("s_permissions", in.SPermissions).
		Where("server_id = ? AND user_id = ?", in.ServerId, in.UserId).
		Update()
	if err != nil {
		return nil, errResponse.DbOperationError("更新失败")
	}
	if err = cache.DelServerUsersCache(ctx, in.ServerId); err != nil {
		return nil, errResponse.OperationFailed("清理缓存失败")
	}
	return &v1.UserRoleRes{
		UserInput: in,
	}, nil
}
