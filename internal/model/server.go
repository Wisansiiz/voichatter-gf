package model

import "github.com/gogf/gf/v2/os/gtime"

type ServerCreateInput struct {
	ServerName   string
	ServerType   string
	ServerImgUrl string
}

type Server struct {
	ServerId      uint64      `json:"serverId"       `
	ServerName    string      `json:"serverName"     `
	CreatorUserId uint64      `json:"creatorUserId" `
	CreateDate    *gtime.Time `json:"createDate"     `
	ServerType    string      `json:"serverType"     `
	ServerImgUrl  string      `json:"serverImgUrl"  `
}
