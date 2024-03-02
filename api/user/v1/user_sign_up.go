package v1

import "github.com/gogf/gf/v2/frame/g"

type SignUpReq struct {
	g.Meta        `path:"/user/sign-up" method:"post" tags:"UserService" summary:"Sign up a new user account"`
	Username      string `v:"required|length:4,16"`
	Email         string `v:"required"`
	PasswordHash  string `v:"required|length:6,16"`
	PasswordHash2 string `v:"required|length:6,16|same:Password"`
}

type SignUpRes struct{}
