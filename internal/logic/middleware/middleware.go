package middleware

import (
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"strings"
	"voichatter/internal/logic/jwt"
	"voichatter/internal/model"
	"voichatter/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Ctx injects custom business context variable into context of current request.
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		Session: r.Session,
	}
	service.BizCtx().Init(r, customCtx)
	if user := service.Session().GetUser(r.Context()); user != nil {
		customCtx.User = &model.ContextUser{
			UserId:   user.UserId,
			Username: user.Username,
		}
	}
	// Continue execution of next middleware.
	r.Middleware.Next()
}

// CORS allows Cross-origin resource sharing.
func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

// Auth 验证是否登录
func (s *sMiddleware) Auth(r *ghttp.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		err := gerror.New("请求头中auth为空")
		r.SetError(err)
		r.Exit()
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if !(len(parts) == 2 && parts[0] == "Bearer") {
		err := gerror.New("请求头中auth格式有误")
		r.SetError(err)
		r.Exit()
	}

	_, err := jwt.ParseToken(parts[1])
	if err != nil {
		err = gerror.New("无效的Token")
		r.SetError(err)
		r.Exit()
	}

	token := parts[1]
	var ctx = gctx.New()
	result, err := g.Redis().Get(ctx, token)
	if err != nil {
		err = gerror.New("redis出错")
		r.SetError(err)
		r.Exit()
	}

	if !result.IsNil() {
		err = gerror.New("Token在黑名单中")
		r.SetError(err)
		r.Exit()
	}

	r.Middleware.Next()
}
