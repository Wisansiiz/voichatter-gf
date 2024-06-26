package model

import "github.com/gogf/gf/v2/os/gtime"

type ServerCreateInput struct {
	ServerName   string
	ServerType   string
	ServerImgUrl string
}

type ServerInfoUpdInput struct {
	ServerId          uint64 `json:"serverId"          `
	ServerName        string `json:"serverName"        `
	ServerType        string `json:"serverType"        `
	ServerImgUrl      string `json:"serverImgUrl"      `
	ServerDescription string `json:"serverDescription" `
}

type Server struct {
	ServerId          uint64      `json:"serverId"          `
	ServerName        string      `json:"serverName"        `
	CreatorUserId     uint64      `json:"creatorUserId"     `
	CreateDate        *gtime.Time `json:"createDate"        `
	ServerType        string      `json:"serverType"        `
	ServerImgUrl      string      `json:"serverImgUrl"      `
	ServerDescription string      `json:"serverDescription" `
}

type ServerPagesInput struct {
	ServerName string `json:"serverName"`
	Page       int    `json:"page"      `
	PageSize   int    `json:"pageSize"  `
}

type ServerPages struct {
	ServerId          uint64      `json:"serverId"          `
	ServerName        string      `json:"serverName"        `
	CreatorUserId     uint64      `json:"creatorUserId"     `
	CreateDate        *gtime.Time `json:"createDate"        `
	ServerType        string      `json:"serverType"        `
	ServerImgUrl      string      `json:"serverImgUrl"      `
	ServerDescription string      `json:"serverDescription" `
}
