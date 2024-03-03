package model

type ChannelCreateInput struct {
	ServerId    uint64
	ChannelName string
	Type        string
}
