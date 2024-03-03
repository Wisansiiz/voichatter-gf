package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta            `path:"/api/register" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Username          string `p:"username"          v:"required|length:4,30#请输入账号|账号长度为:{min}到:{max}位"`
	Email             string `p:"email"             v:"required"`
	PasswordHash      string `p:"password"          v:"required|length:6,30#请输入密码|密码长度不够"`
	ReenteredPassword string `p:"reenteredPassword" v:"required|length:6,30|same:password#请确认密码|密码长度不够|两次密码不一致"`
}

type SignUpRes struct{}
