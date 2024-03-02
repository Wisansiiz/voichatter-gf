package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SignInReq struct {
	g.Meta       `path:"/api/login" method:"post" tags:"UserService" summary:"Sign in with exist account"`
	Username     string `p:"username" v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	PasswordHash string `p:"password" v:"required|length:6,30#请输入密码|密码长度不够"`
}
type SignInRes struct {
	Token string `json:"token"`
}
