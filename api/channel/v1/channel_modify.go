package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ChannelModifyReq struct {
	g.Meta      `path:"/channel/{serverId}/{channelId}" method:"put" tags:"ChannelService" summary:"修改分组名称"`
	ServerId    uint64 `p:"serverId" v:"required"`
	ChannelId   uint64 `p:"channelId" v:"required"`
	ChannelName string `p:"channelName" v:"required|length:4,10#请输入频道名|频道名长度为:{min}到:{max}位"`
}

type ChannelModifyRes struct {
	Channel *model.ChannelInfo `json:"channel"`
}
