package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type GroupModifyReq struct {
	g.Meta    `path:"/group/{serverId}/{groupId}" method:"put" tags:"GroupService" summary:"修改分组名称"`
	ServerId  uint64 `p:"serverId" v:"required"`
	GroupId   uint64 `p:"groupId" v:"required"`
	GroupName string `p:"groupName" v:"required|length:4,10#请输入分组名|分组名长度为:{min}到:{max}位"`
}

type GroupModifyRes struct {
	Group *model.Group `json:"group"`
}
