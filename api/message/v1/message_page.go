package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type MessagePagesReq struct {
	g.Meta    `path:"/messages/page/{serverId}/{channelId}" method:"get" tags:"MessageService" summary:"message-pages"`
	ServerId  uint64 `p:"serverId" v:"required"`
	ChannelId uint64 `p:"channelId" v:"required"`
	Page      int    `json:"page" example:"10" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PageSize  int    `json:"pageSize" example:"1" d:"10" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}

type MessagePagesRes struct {
	MessagePages *model.MessagePagesRes `json:"messagePages"`
}
