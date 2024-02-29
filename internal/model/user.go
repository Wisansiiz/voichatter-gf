package model

import "time"

type UserCreateInput struct {
	Username     string `json:"username"`
	Email        string `json:"email"`
	PasswordHash string `json:"password"`
}

type UserSignInInput struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password"`
}

type UserList4Server struct {
	UserID        uint       `json:"user_id"`
	Username      string     `json:"username"`
	Email         string     `json:"email"`
	AvatarURL     string     `json:"avatar_url"`
	SPermissions  string     `json:"s_permissions"`
	LastLoginDate *time.Time `json:"last_login_date"`
}
