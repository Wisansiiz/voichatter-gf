package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type UserInfoUpdReq struct {
	g.Meta   `path:"/user/info" method:"put" tags:"UserService" summary:"修改用户信息"`
	Username string `p:"username"          v:"length:4,10#账号长度为:{min}到:{max}位"`
	Email    string `p:"email"             v:"email"`
}

type UserInfoUpdRes struct {
	UserInfo *model.UserInfo `json:"userInfo"`
}
