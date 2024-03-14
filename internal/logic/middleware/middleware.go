package middleware

import (
	"encoding/json"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"voichatter/internal/dao"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
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

type Msg struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

// WebSocketAuth websocket鉴权
func (s *sMiddleware) WebSocketAuth(gfToken *gtoken.GfToken, r *ghttp.Request, conn *websocket.Conn) {
	var currentUserID string
	// 接收和处理消息
	var msg Msg
	_, p, err := conn.ReadMessage()
	if err != nil {
		return
	}
	if err = json.Unmarshal(p, &msg); err != nil {
		return
	}
	if msg.Code == "authorization" {
		token := msg.Data["token"]
		rep := gfToken.DecryptToken(r.Context(), token.(string)).Data
		m := gconv.Map(rep)
		key := m["userKey"]
		if key == nil {
			r.SetError(gerror.New("未登录"))
			r.Exit()
		}
		jsonBytes, _ := json.Marshal(g.Map{
			"code": "authorization",
			"data": g.Map{
				"userId": key,
			},
		})
		if err = conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
			r.SetError(gerror.New("出错了"))
			r.Exit()
		}
		currentUserID = gconv.String(key)
	}
	var user entity.User
	err = dao.User.Ctx(r.Context()).Where("user_id = ?", gconv.Uint64(currentUserID)).Scan(&user)
	if err != nil {
		return
	}
	r.SetCtxVar("userId", currentUserID)
	r.SetCtxVar("email", user.Email)
	r.SetCtxVar("avatarUrl", user.AvatarUrl)
	r.SetCtxVar("username", user.Username)
	r.Middleware.Next()
}
