package chat

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
)

var (
	clients       = make(map[*websocket.Conn]string)
	groupChannels = make(map[string][]*websocket.Conn)
	conn          *websocket.Conn
)

type Msg struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

func GroupChat(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.SetError(gerror.New("websocket error"))
		r.Exit()
	}
	conn = ws.Conn
	defer r.Exit()

	currentUserId := r.GetCtxVar("userId").String()
	channelId := r.GetQuery("serverId").String() + r.GetQuery("channelId").String()
	clients[conn] = currentUserId
	// 接收和处理消息
	for {
		var msg Msg
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err = gconv.Struct(p, &msg); err != nil {
			break
		}
		if msg.Code == "offer" {
			targetId := msg.Data["targetId"]
			offer := msg.Data["offer"]
			broadcastRTCMessage(targetId, currentUserId, "offer", offer)
		} else if msg.Code == "answer" {
			targetId := msg.Data["targetId"]
			answer := msg.Data["answer"]
			broadcastRTCMessage(targetId, currentUserId, "answer", answer)
		} else if msg.Code == "icecandidate" {
			targetId := msg.Data["targetId"]
			candidate := msg.Data["candidate"]
			broadcastRTCMessage(targetId, currentUserId, "icecandidate", candidate)
		} else if msg.Code == "join_group" {
			groupChannels[channelId] = append(groupChannels[channelId], conn)
			if broadcastGroups(msg.Code, channelId, currentUserId) {
				return
			}
		} else if msg.Code == "leave_group" {
			if broadcastGroups(msg.Code, channelId, currentUserId) {
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

func broadcastGroups(code string, channelId string, currentUserID string) bool {
	connections := groupChannels[channelId]
	for _, numbers := range connections {
		if conn != numbers {
			message := g.Map{
				"code": code,
				"data": g.Map{
					"fromId": currentUserID,
				},
			}
			jsonBytes, _ := gjson.Marshal(message)
			if err := numbers.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				return true
			}
		}
	}
	return false
}

func broadcastRTCMessage(targetId any, currentUserID string, code string, data any) {
	for clientConn, userId := range clients {
		if userId == targetId.(string) {
			message := g.Map{
				"code": code,
				"data": g.Map{
					"fromId": currentUserID,
					code:     data,
				},
			}
			jsonBytes, _ := gjson.Marshal(message)
			if err := clientConn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				return
			}
		}
	}
}
