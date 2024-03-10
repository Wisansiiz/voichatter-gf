package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type UserIdReq struct {
	g.Meta `path:"/user" method:"get" tags:"UserService" summary:""`
}

type UserIdRes struct {
	UserId uint64 `json:"userId"`
}
