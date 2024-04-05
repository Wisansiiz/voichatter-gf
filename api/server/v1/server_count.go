package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ServerCountReq struct {
	g.Meta `path:"/servers/count" method:"get" tags:"ServerService" summary:""`
}
type ServerCountRes struct {
	Count uint64 `json:"count"`
}
