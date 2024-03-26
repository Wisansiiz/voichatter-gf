package model

type NotificationCreateInput struct {
	Title    string `json:"title"         ` // 标题
	Content  string `json:"content"       ` // 内容
	ServerId uint64 `json:"serverId"      ` // 属于哪个服务器
}

type Notification struct {
	NotificationId uint64 `json:"notificationId"` // 通知id
	Title          string `json:"title"         ` // 标题
	Content        string `json:"content"       ` // 内容
	ServerId       uint64 `json:"serverId"      ` // 属于哪个服务器
	CreateUserId   uint64 `json:"createUserId"  ` // 创建者id
}
