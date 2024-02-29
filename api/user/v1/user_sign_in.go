package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SignInReq struct {
	g.Meta       `path:"/user/sign-in" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Username     string `v:"required"`
	PasswordHash string `v:"required"`
}
type SignInRes struct {
	Token string `json:"token"`
}
