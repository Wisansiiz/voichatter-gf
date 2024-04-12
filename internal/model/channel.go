package model

type ChannelCreateInput struct {
	ServerId    uint64
	GroupId     *uint64
	ChannelName string
	Type        string
}

type ChannelInfo struct {
	ChannelId    uint64 `json:"channelId"     `  // 频道id
	ChannelName  string `json:"channelName"   `  // 频道名称
	ServerId     uint64 `json:"serverId"      `  // 服务器id
	GroupId      uint64 `json:"groupId"       `  // 群组id
	Type         string `json:"type"           ` // 服务器类型
	CreateUserId uint64 `json:"createUserId" `   // 服务器创建者id
}

type ChannelModifyInput struct {
	ServerId    uint64 `json:"serverId"`
	ChannelId   uint64 `json:"channelId"`
	ChannelName string `json:"channelName"`
}

type ChannelRemoveInput struct {
	ServerId  uint64 `json:"serverId"`
	ChannelId uint64 `json:"channelId"`
}
