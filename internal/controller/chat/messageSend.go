package chat

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"github.com/importcjj/sensitive"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
)

var (
	rooms  = make(map[uint64][]map[uint64]*websocket.Conn)
	filter = sensitive.New()
)

func init() {
	if err := filter.LoadWordDict("./utility/dict/dict.txt"); err != nil {
		return
	}
}

func MessageSend(r *ghttp.Request) {
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

	userId := r.GetCtxVar("userId").Uint64()
	g.Dump(userId)
	// 接收和处理消息
	var targetId uint64
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		str, _ := gjson.DecodeToJson(p)
		targetId = str.Get("targetId").Uint64()
		for i, m := range rooms[targetId] {
			if c, ok := m[userId]; ok {
				if c != conn {
					delete(m, userId)
					m[userId] = conn
					break
				} else {
					break
				}
			}
			if i == len(rooms[targetId])-1 {
				currentUser := make(map[uint64]*websocket.Conn)
				currentUser[userId] = conn
				rooms[targetId] = append(rooms[targetId], currentUser)
			}
		}
		if len(rooms[targetId]) == 0 {
			currentUser := make(map[uint64]*websocket.Conn)
			currentUser[userId] = conn
			rooms[targetId] = append(rooms[targetId], currentUser)
		}
		code := str.Get("code").String()
		data := str.Get("data").String()

		content := filter.Replace(data, '*')
		var id int64
		if code != "ping" && code != "update" {
			// 持久化消息到数据库
			id, err = dao.Message.Ctx(r.Context()).InsertAndGetId(
				&entity.Message{
					MessageType:  code,
					SenderUserId: userId,
					Content:      content,
					ChannelId:    targetId,
					ServerId:     str.Get("serverId").Uint64(),
					SendDate:     gtime.Now(),
				},
			)
			if err != nil {
				r.SetError(gerror.New("持久化出错"))
				return
			}
		}
		// 根据收到的code值修改消息的内容
		// when code is "update", update the message content
		if code == "update" {
			content = content + " (edited)"
			id = str.Get("messageId").Int64()
			_, err = dao.Message.Ctx(r.Context()).
				Fields("content").
				Data(g.Map{"content": content}).
				Where("message_id", id).
				Update()
			if err != nil {
				r.SetError(gerror.New("更新消息出错"))
				return
			}
		}
		var messageInfo = model.MessageInfo{
			MessageId:    uint64(id),
			MessageType:  code,
			ChannelId:    targetId,
			Content:      content,
			ServerId:     str.Get("serverId").Uint64(),
			Attachment:   "",
			AvatarUrl:    r.GetCtxVar("avatarUrl").String(),
			SenderUserId: userId,
			Username:     r.GetCtxVar("username").String(),
			SendDate:     gtime.Now(),
		}
		// 将消息广播给房间内的所有客户端
		go broadcastMessage(targetId, gconv.Bytes(
			g.Map{
				"count":   len(rooms[targetId]),
				"code":    code,
				"message": messageInfo,
			},
		))
	}
	defer clean(targetId, userId)
}

func clean(channelId uint64, userId uint64) {
	mu.Lock()
	defer mu.Unlock()
	// 在连接关闭时，将其从房间中移除
	connections := rooms[channelId]
	for i, number := range connections {
		if _, ok := number[userId]; ok {
			delete(number, userId)
			rooms[channelId] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func broadcastMessage(channelId uint64, message []byte) {
	mu.RLock()
	defer mu.RUnlock()
	// 查询房间内所有连接的客户端
	connections := rooms[channelId]
	for _, conn := range connections {
		// 发送消息给所有客户端包括自己
		for _, v := range conn {
			if err := v.WriteMessage(websocket.TextMessage, message); err != nil {
				return
			}
		}
	}
}
