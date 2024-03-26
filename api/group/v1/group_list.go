package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type GroupListReq struct {
	g.Meta   `path:"/groups/{serverId}" method:"get" tags:"GroupService" summary:"查询分组列表"`
	ServerId uint64 `p:"serverId" v:"required"`
}

type GroupListRes struct {
	ChannelList []*model.ChannelInfo `json:"channelList"`
	GroupList   []*model.GroupList   `json:"groupList"`
}
