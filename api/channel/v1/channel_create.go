package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ChannelCreateReq struct {
	g.Meta      `path:"/channel/{serverId}" method:"post" tags:"ChannelService" summary:""`
	ServerId    uint64 `p:"serverId" v:"required"`
	ChannelName string `p:"channelName" v:"required"`
	Type        string `p:"type" v:"required"`
}

type ChannelCreateRes struct {
	Channel *model.ChannelInfo `json:"channel"`
}
