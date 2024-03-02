package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"voichatter/internal/controller/chat"
	"voichatter/internal/controller/hello"
	"voichatter/internal/controller/user"
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
					service.Middleware().Ctx,
					ghttp.MiddlewareHandlerResponse,
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.POST("/", user.NewV1().SignIn)
					group.POST("/", user.NewV1().SignUp)
					//group.GET("/", user.NewV1().ServerList)
				})
				group.GET("/api/yy", func(r *ghttp.Request) {
					chat.HandleWebSocket(r)
				})
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(service.Middleware().Auth)
					group.Bind(
						hello.NewV1(),
						user.NewV1().SignOut,
						user.NewV1().ServerList,
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
