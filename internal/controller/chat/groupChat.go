package chat

import (
	"encoding/json"
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"voichatter/internal/service"
)

var (
	clients       = make(map[*websocket.Conn]string)
	groupChannels = make(map[string][]*websocket.Conn)
)

type Msg struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

func GroupChat(gfToken *gtoken.GfToken, r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.SetError(gerror.New("websocket error"))
		r.Exit()
	}
	conn := ws.Conn
	defer r.Exit()

	// 验证身份
	service.Middleware().WebSocketAuth(gfToken, r, conn)
	currentUserId := r.GetCtxVar("userId").String()
	channelId := r.GetQuery("channelId").String()
	clients[conn] = currentUserId
	// 接收和处理消息
	for {
		var msg Msg
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err = json.Unmarshal(p, &msg); err != nil {
			break
		}
		if msg.Code == "offer" {
			targetId := msg.Data["targetId"]
			offer := msg.Data["offer"]
			broadcastRTCMessage(targetId, currentUserId, "offer", "offer", offer)
		} else if msg.Code == "answer" {
			targetId := msg.Data["targetId"]
			answer := msg.Data["answer"]
			broadcastRTCMessage(targetId, currentUserId, "answer", "answer", answer)
		} else if msg.Code == "icecandidate" {
			targetId := msg.Data["targetId"]
			candidate := msg.Data["candidate"]
			broadcastRTCMessage(targetId, currentUserId, "icecandidate", "candidate", candidate)
		} else if msg.Code == "join_group" {
			groupChannels[channelId] = append(groupChannels[channelId], conn)
			if broadcastGroups(msg.Code, channelId, conn, currentUserId) {
				return
			}
		} else if msg.Code == "leave_group" {
			if broadcastGroups(msg.Code, channelId, conn, currentUserId) {
				return
			}
			break
		}
	}

	if channelId != "" {
		// 在连接关闭时，将其从房间中移除
		connections := groupChannels[channelId]
		for i, numbers := range connections {
			if numbers == conn {
				groupChannels[channelId] = append(connections[:i], connections[i+1:]...)
				break
			}
		}
	}
	// 在连接关闭时，将其从连接中移除
	for clientConn := range clients {
		if clientConn == conn {
			delete(clients, clientConn)
		}
	}
}

func broadcastGroups(code string, channelId string, conn *websocket.Conn, currentUserID string) bool {
	connections := groupChannels[channelId]
	for _, numbers := range connections {
		if conn != numbers {
			message := g.Map{
				"code": code,
				"data": g.Map{
					"fromId": currentUserID,
				},
			}
			jsonBytes, _ := json.Marshal(message)
			if err := numbers.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				return true
			}
		}
	}
	return false
}

func broadcastRTCMessage(targetId any, currentUserID string, code string, dataName string, data any) {
	for clientConn, userId := range clients {
		if userId == targetId.(string) {
			message := g.Map{
				"code": code,
				"data": g.Map{
					"fromId": currentUserID,
					dataName: data,
				},
			}
			jsonBytes, _ := json.Marshal(message)
			if err := clientConn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				return
			}
		}
	}
}