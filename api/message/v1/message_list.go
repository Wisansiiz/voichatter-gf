package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type MessageListReq struct {
	g.Meta    `path:"/messages/{serverId}/{channelId}" method:"get" tags:"MessageService" summary:"message-list"`
	ServerId  uint64 `p:"serverId" v:"required"`
	ChannelId uint64 `p:"channelId" v:"required"`
}

type MessageListRes struct {
	MessageList []*model.MessageInfo `json:"messageList"`
}
