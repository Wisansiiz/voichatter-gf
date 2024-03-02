package model

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type Context struct {
	Session *ghttp.Session // Session in context.
	User    *ContextUser   // User in context.
}

type ContextUser struct {
	UserId   uint64 `json:"user_id"`  // User ID.
	Username string `json:"username"` // User nickname.
}
