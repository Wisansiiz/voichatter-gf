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
	clients       = make(map[string]*websocket.Conn)
	groupChannels = make(map[string][]*websocket.Conn)
	mu            sync.RWMutex
)

func GroupChat(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.SetError(gerror.New("websocket error"))
		r.Exit()
	}
	conn := ws.Conn
	defer func(ws *ghttp.WebSocket) {
		err = ws.Close()
		if err != nil {
			return
		}
	}(ws)

	currentUserId := r.GetCtxVar("userId").String()
	channelId := r.GetQuery("serverId").String() + r.GetQuery("channelId").String()
	addClient(currentUserId, conn)

	for {
		msg := Msg{}
		if err := readMessage(conn, &msg); err != nil {
			break
		}
		if ok := handleMessage(channelId, currentUserId, conn, &msg); ok {
			break
		}
	}

	cleanUp(channelId, currentUserId, conn)
}

func addClient(userId string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	clients[userId] = conn
}

func readMessage(conn *websocket.Conn, msg *Msg) error {
	_, p, err := conn.ReadMessage()
	if err != nil {
		return err
	}
	return gjson.Unmarshal(p, msg)
}

func handleMessage(channelId, currentUserId string, conn *websocket.Conn, msg *Msg) bool {
	switch msg.Code {
	case "offer", "answer", "candidate":
		targetId := msg.Data["targetId"]
		broadcastRTCMessage(targetId, currentUserId, msg)
	case "join_group":
		addToGroup(channelId, conn)
		broadcastGroupMessage(channelId, currentUserId, msg)
	case "leave_group":
		broadcastGroupMessage(channelId, currentUserId, msg)
		removeFromGroup(channelId, conn)
		return true
	}
	return false
}

func cleanUp(channelId, userId string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	delete(clients, userId)
	if channelId != "" {
		connections := groupChannels[channelId]
		for i, c := range connections {
			if c == conn {
				groupChannels[channelId] = append(connections[:i], connections[i+1:]...)
				break
			}
		}
	}
}

func addToGroup(channelId string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	groupChannels[channelId] = append(groupChannels[channelId], conn)
}

func removeFromGroup(channelId string, conn *websocket.Conn) {
	mu.Lock()
	defer mu.Unlock()
	connections := groupChannels[channelId]
	for i, c := range connections {
		if c == conn {
			groupChannels[channelId] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func broadcastRTCMessage(targetId any, currentUserID string, msg *Msg) {
	mu.RLock()
	defer mu.RUnlock()
	targetConn, ok := clients[gconv.String(targetId)]
	if !ok {
		return
	}
	sendMessage(targetConn, msg.Code, g.Map{
		"fromId": currentUserID,
		msg.Code: msg.Data[msg.Code],
	})
}

func broadcastGroupMessage(channelId, currentUserID string, msg *Msg) {
	mu.RLock()
	defer mu.RUnlock()
	connections := groupChannels[channelId]
	for _, conn := range connections {
		if clients[currentUserID] != conn {
			sendMessage(conn, msg.Code, g.Map{
				"fromId": currentUserID,
			})
		}
	}
}

func sendMessage(conn *websocket.Conn, code string, data g.Map) {
	message := &Msg{
		Code: code,
		Data: data,
	}
	jsonBytes, _ := gjson.Marshal(message)
	err := conn.WriteMessage(websocket.TextMessage, jsonBytes)
	if err != nil {
		return
	}
}
