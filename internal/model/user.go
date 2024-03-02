package model

import "time"

type UserCreateInput struct {
	Username     string
	Email        string
	PasswordHash string
}

type UserSignInInput struct {
	Username     string
	PasswordHash string
}

type UserList4Server struct {
	UserID        uint
	Username      string
	Email         string
	AvatarURL     string
	SPermissions  string
	LastLoginDate *time.Time
}
