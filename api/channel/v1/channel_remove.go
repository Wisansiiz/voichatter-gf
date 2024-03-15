package v1

import "github.com/gogf/gf/v2/frame/g"

type ChannelRemoveReq struct {
	g.Meta    `path:"/channel/{serverId}/{channelId}" method:"delete" tags:"ChannelService" summary:"删除分组"`
	ServerId  uint64 `p:"serverId" v:"required"`
	ChannelId uint64 `p:"channelId" v:"required"`
}

type ChannelRemoveRes struct{}
