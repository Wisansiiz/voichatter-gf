// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/server/v1"
	"voichatter/internal/model"
)

type (
	IServer interface {
		ServerList(ctx context.Context, _ *v1.ServerListReq) (res *v1.ServerListRes, err error)
		ServerCreate(ctx context.Context, in model.ServerCreateInput) (res *v1.ServerCreateRes, err error)
		ServerJoin(ctx context.Context, serverId uint64) (res *v1.ServerJoinRes, err error)
		ServerDel(ctx context.Context, serverId uint64) (res *v1.ServerDelRes, err error)
	}
)

var (
	localServer IServer
)

func Server() IServer {
	if localServer == nil {
		panic("implement not found for interface IServer, forgot register?")
	}
	return localServer
}

func RegisterServer(i IServer) {
	localServer = i
}
