package model

type GroupList struct {
	GroupId     uint64        `json:"groupId"   `  // 分组id
	ServerId    uint64        `json:"serverId"  `  // 服务器id
	GroupName   string        `json:"groupName" `  // 分组名称
	ChannelList []ChannelInfo `json:"channelList"` // 频道列表
}