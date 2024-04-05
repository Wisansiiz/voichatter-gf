package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta            `path:"/register" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Username          string `p:"username"          v:"required|length:4,10#请输入账号|账号长度为:{min}到:{max}位"`
	Email             string `p:"email"             v:"email"`
	PasswordHash      string `p:"password"          v:"password|length:6,18#请输入密码|密码长度不够"`
	ReenteredPassword string `p:"reenteredPassword" v:"password|length:6,18|same:password#请确认密码|密码长度不够|两次密码不一致"`
	Code              string `p:"code"              v:"required#请输入验证码"`
	Id                string `p:"id"                v:"required"`
}

type SignUpRes struct{}
