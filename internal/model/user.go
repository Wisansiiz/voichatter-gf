package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type UserCreateInput struct {
	Username          string
	Email             string
	PasswordHash      string
	ReenteredPassword string
	Code              string
	Id                string
}

type UserSignInInput struct {
	g.Meta       `path:"/login" method:"post" tags:"" summary:"Sign in with exist account"`
	Username     string `p:"username" v:"required|length:4,20#请输入账号|账号长度为:{min}到:{max}位"`
	PasswordHash string `p:"password" v:"required|length:6,20#请输入密码|密码长度不够"`
}

type UserList4Server struct {
	UserID        uint        `json:"userID"`
	Username      string      `json:"username"`
	Email         string      `json:"email"`
	AvatarURL     string      `json:"avatarURL"`
	SPermissions  string      `json:"SPermissions"`
	LastLoginDate *gtime.Time `json:"lastLoginDate"`
}

type LoginRes struct {
	UserId        uint64
	Username      string
	Email         string
	LastLoginDate *gtime.Time
	AvatarUrl     string
}

type ModifyUserRoleInput struct {
	UserId       uint64 `json:"userId"`
	SPermissions string `json:"SPermissions"`
	ServerId     uint64 `json:"serverId"`
}

type UserInfo struct {
	UserId    uint64 `json:"userId"           ` // 用户id
	Username  string `json:"username"         ` // 用户名
	Email     string `json:"email"            ` // 邮箱
	AvatarUrl string `json:"avatarUrl"        ` // 头像链接
}
