package v1

import "github.com/gogf/gf/v2/frame/g"

type SignOutReq struct {
	g.Meta `path:"/api/logout" method:"post" tags:"UserService" summary:"Sign out current user"`
}
type SignOutRes struct{}
