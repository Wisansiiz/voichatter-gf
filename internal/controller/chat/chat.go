package chat

import (
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gorilla/websocket"
	"log"
)

var (
	clients       = make(map[*websocket.Conn]string)
	groupChannels = make(map[string][]*websocket.Conn)
)

type Msg struct {
	Code string         `json:"code"`
	Data map[string]any `json:"data"`
}

func HandleWebSocket(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.Exit()
	}
	conn := ws.Conn
	defer r.Exit()

	channelId := r.GetQuery("channelId").String()
	currentUserID := r.GetQuery("id").String()
	clients[conn] = currentUserID
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
			broadcastMessage(targetId, currentUserID, "offer", "offer", offer)
		} else if msg.Code == "answer" {
			targetId := msg.Data["targetId"]
			answer := msg.Data["answer"]
			broadcastMessage(targetId, currentUserID, "answer", "answer", answer)
		} else if msg.Code == "icecandidate" {
			targetId := msg.Data["targetId"]
			candidate := msg.Data["candidate"]
			broadcastMessage(targetId, currentUserID, "icecandidate", "candidate", candidate)
		} else if msg.Code == "join_group" {
			groupChannels[channelId] = append(groupChannels[channelId], conn)
			if broadcastGroups(msg.Code, channelId, conn, currentUserID) {
				return
			}
		} else if msg.Code == "leave_group" {
			if broadcastGroups(msg.Code, channelId, conn, currentUserID) {
				return
			}
			break
		} else {
			jsonBytes, _ := json.Marshal(msg)
			if err := conn.WriteMessage(websocket.TextMessage, jsonBytes); err != nil {
				log.Println("出错了")
				r.Exit()
			}
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

func broadcastMessage(targetId any, currentUserID string, code string, dataName string, data any) {
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
