package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type UserRoleReq struct {
	g.Meta       `path:"/user/{serverId}" method:"put" tags:"UserService" summary:""`
	ServerId     uint64 `p:"serverId" v:"required"`
	UserId       uint64 `p:"userId" v:"required"`
	SPermissions string `p:"sPermissions" v:"required"`
}

type UserRoleRes struct {
	UserInput model.ModifyUserRoleInput `json:"userInput"`
}
