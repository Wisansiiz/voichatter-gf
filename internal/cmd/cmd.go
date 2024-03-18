package cmd

import (
	"context"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"voichatter/internal/controller/channel"
	"voichatter/internal/controller/chat"
	sgroup "voichatter/internal/controller/group"
	"voichatter/internal/controller/message"
	"voichatter/internal/controller/server"
	"voichatter/internal/controller/user"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
)

const (
	swaggerUIPageContent = `
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
  <meta name="description" content="SwaggerUI"/>
  <title>SwaggerUI</title>
  <link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist@latest/swagger-ui.css" />
</head>
<body>
<div id="swagger-ui"></div>
<script src="https://unpkg.com/swagger-ui-dist@latest/swagger-ui-bundle.js" crossorigin></script>
<script>
	window.onload = () => {
		window.ui = SwaggerUIBundle({
			url:    '/api.json',
			dom_id: '#swagger-ui',
		});
	};
</script>
</body>
</html>
`
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.GET("/swagger", func(r *ghttp.Request) {
					r.Response.Write(swaggerUIPageContent)
				})
				group.Middleware(
					ghttp.MiddlewareCORS, //跨域中间件
					ghttp.MiddlewareHandlerResponse,
				)
				// 启动gToken
				gfToken := &gtoken.GfToken{
					LoginPath:        "/login",
					CacheMode:        2,
					LoginBeforeFunc:  service.User().LoginFunc,
					LogoutPath:       "/logout",
					AuthAfterFunc:    authFunc,
					LoginAfterFunc:   loginAfterFunc,
					MultiLogin:       true,
					AuthExcludePaths: g.SliceStr{"/register, /login"},
				}
				group.Group("/", func(group *ghttp.RouterGroup) {
					err := gfToken.Middleware(ctx, group)
					if err != nil {
						panic(err)
					}
					group.Group("/api", func(group *ghttp.RouterGroup) {
						group.GET("/wz", chat.MessageSend)
						group.GET("/yy", chat.GroupChat)
					})
					group.Bind(
						sgroup.NewV1(),
						user.NewV1(),
						server.NewV1(),
						message.NewV1(),
						channel.NewV1(),
					)
				})
			})
			s.EnableHTTPS("./cert/server.pem", "./cert/server-key.pem")
			s.SetPort(9000)
			s.Run()
			return nil
		},
	}
)

func authFunc(r *ghttp.Request, respData gtoken.Resp) {
	var userInfo model.LoginRes
	respData.Code = gtoken.SUCCESS
	var s gTokenData
	err := gconv.Struct(respData.Data, &s)
	if err != nil {
		r.SetError(gerror.New("验证失败"))
		r.Exit()
	}
	userInfo = model.LoginRes{
		UserId:        uint64(s.Data.UserId),
		Username:      s.Data.Username,
		Email:         s.Data.Email,
		AvatarUrl:     s.Data.AvatarUrl,
		LastLoginDate: &s.Data.LastLoginDate,
	}
	r.SetCtxVar("userId", userInfo.UserId)
	r.SetCtxVar("username", userInfo.Username)
	r.SetCtxVar("avatarUrl", userInfo.AvatarUrl)
	r.Middleware.Next()
}

type gTokenData struct {
	CreateTime int64 `json:"createTime"`
	Data       struct {
		UserId           int         `json:"user_id"`
		Username         string      `json:"username"`
		Email            string      `json:"email"`
		PasswordHash     string      `json:"password_hash"`
		AvatarUrl        string      `json:"avatar_url"`
		RegistrationDate gtime.Time  `json:"registration_date"`
		LastLoginDate    gtime.Time  `json:"last_login_date"`
		DeletedAt        interface{} `json:"deleted_at"`
	} `json:"data"`
	RefreshTime int64  `json:"refreshTime"`
	UserKey     string `json:"userKey"`
	Uuid        string `json:"uuid"`
}

func loginAfterFunc(r *ghttp.Request, respData gtoken.Resp) {
	type data struct {
		Uuid    string `json:"uuid"`
		Token   string `json:"token"`
		UserKey string `json:"userKey"`
	}
	var dataRep data
	err := gconv.Struct(respData.Data, &dataRep)
	if err != nil {
		r.SetError(gerror.New("数据错误"))
		r.Exit()
	}
	userId := gconv.Uint64(dataRep.UserKey)
	var userInfo model.LoginRes
	err = g.Model(entity.User{}).Where("user_id = ?", userId).Scan(&userInfo)
	if err != nil {
		r.SetError(gerror.New("数据错误"))
		r.Exit()
	}
	r.Response.WriteJson(g.Map{
		"code":    0,
		"message": "登录成功",
		"data": g.Map{
			"token":     dataRep.Token,
			"userId":    userInfo.UserId,
			"username":  userInfo.Username,
			"avatarUrl": userInfo.AvatarUrl,
		},
	})
	r.Middleware.Next()
}
