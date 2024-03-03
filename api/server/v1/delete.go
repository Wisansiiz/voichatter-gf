package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ServerDelReq struct {
	g.Meta `path:"/servers" method:"delete" tags:"ServerService" summary:"servers"`
}

type ServerDelRes struct{}
