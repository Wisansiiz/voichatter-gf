package chat

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gorilla/websocket"
	"voichatter/internal/dao"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
)

var (
	rooms = make(map[uint64][]*websocket.Conn)
)

func MessageSend(gfToken *gtoken.GfToken, r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.SetError(gerror.New("websocket error"))
		r.Exit()
	}
	conn := ws.Conn
	defer r.Exit()

	// 验证身份
	service.Middleware().WebSocketAuth(gfToken, r, conn)
	userId := r.GetCtxVar("userId").Uint64()
	// 从URL参数获取频道号
	serverId := r.GetQuery("serverId").Uint64()
	channelId := r.GetQuery("channelId").Uint64()
	// 将连接加入到对应的房间
	rooms[channelId] = append(rooms[channelId], conn)
	// 接收和处理消息
	for {
		_, p, err := conn.ReadMessage()
		if err != nil {
			break
		}
		// 持久化消息到数据库
		content := string(p)
		_, err = dao.Message.Ctx(r.Context()).Insert(
			&entity.Message{
				SenderUserId: userId,
				Content:      content,
				ChannelId:    channelId,
				ServerId:     serverId,
				SendDate:     gtime.Now(),
			},
		)
		if err != nil {
			r.SetError(gerror.New("持久化出错"))
			return
		}
		// 将消息广播给房间内的所有客户端
		go broadcastMessage(channelId, p)
	}

	// 在连接关闭时，将其从房间中移除
	connections := rooms[channelId]
	for i, number := range connections {
		if number == conn {
			rooms[channelId] = append(connections[:i], connections[i+1:]...)
			break
		}
	}
}

func broadcastMessage(channelId uint64, message []byte) {
	// 查询房间内所有连接的客户端
	connections := rooms[channelId]
	for _, conn := range connections {
		// 发送消息给所有客户端包括自己
		if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
			return
		}
	}
}