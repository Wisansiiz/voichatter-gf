package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type GroupCreateReq struct {
	g.Meta    `path:"/group/{serverId}" method:"post" tags:"GroupService" summary:"创建服务器分组"`
	ServerId  uint64 `p:"serverId" v:"required"`
	GroupName string `p:"groupName" v:"required|length:4,10#请输入分组名|分组名长度为:{min}到:{max}位"`
}

type GroupCreateRes struct {
	Group *model.Group `json:"group"`
}
