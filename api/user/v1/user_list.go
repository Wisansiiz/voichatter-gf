package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type UserListReq struct {
	g.Meta   `path:"/users/{serverId}" method:"get" tags:"" summary:""`
	ServerId uint64 `p:"serverId" v:"required"`
}

type UserListRes struct {
	Users *[]model.UserList4Server `json:"users"`
}
