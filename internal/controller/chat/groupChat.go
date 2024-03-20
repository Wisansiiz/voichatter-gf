package chat

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"sync"
)

type Msg struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

var (
	clients       = make(map[*websocket.Conn]string)
	groupChannels = make(map[string][]*websocket.Conn)
	msg           Msg
	mu            sync.RWMutex
)

func GroupChat(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.SetError(gerror.New("websocket error"))
		r.Exit()
	}
	conn := ws.Conn
	defer r.Exit()
	defer ws.Close()

	currentUserId := r.GetCtxVar("userId").String()
	channelId := r.GetQuery("serverId").String() + r.GetQuery("channelId").String()
	mu.Lock()
	clients[conn] = currentUserId
	groupChannels[channelId] = append(groupChannels[channelId], conn)
	mu.Unlock()
	// 接收和处理消息
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		if err = gjson.Unmarshal(p, &msg); err != nil {
			break
		}
		if msg.Code == "offer" {
			targetId := msg.Data["targetId"]
			offer := msg.Data["offer"]
			broadcastRTCMessage(targetId, currentUserId, offer)
		} else if msg.Code == "answer" {
			targetId := msg.Data["targetId"]
			answer := msg.Data["answer"]
			broadcastRTCMessage(targetId, currentUserId, answer)
		} else if msg.Code == "candidate" {
			targetId := msg.Data["targetId"]
			candidate := msg.Data["candidate"]
			broadcastRTCMessage(targetId, currentUserId, candidate)
		} else if msg.Code == "join_group" {
			//groupChannels[channelId] = append(groupChannels[channelId], conn)
			if broadcastGroups(channelId, conn, currentUserId) {
				return
			}
		} else if msg.Code == "leave_group" {
			if broadcastGroups(channelId, conn, currentUserId) {
				return
			}
			break
		}
	}

	defer cleanUp(channelId, conn)
}

func cleanUp(channelId string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
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

func broadcastGroups(channelId string, conn *websocket.Conn, currentUserID string) bool {
	mu.RLock()
	connections := groupChannels[channelId]
	mu.RUnlock()
	for _, numbers := range connections {
		if conn != numbers {
			message := &Msg{
				Code: msg.Code,
				Data: g.Map{
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

func broadcastRTCMessage(targetId any, currentUserID string, data any) {
	mu.RLock()
	defer mu.RUnlock()
	for clientConn, userId := range clients {
		if userId == gconv.String(targetId) {
			message := &Msg{
				Code: msg.Code,
				Data: g.Map{
					"fromId": currentUserID,
					msg.Code: data,
				},
			}
			jsonBytes, _ := gjson.Marshal(message)
			if err := clientConn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				return
			}
		}
	}
}
