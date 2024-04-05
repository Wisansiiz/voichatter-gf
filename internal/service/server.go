// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "voichatter/api/server/v1"
	"voichatter/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IServer interface {
		ServerList(ctx context.Context, _ *v1.ServerListReq) (res *v1.ServerListRes, err error)
		ServerCreate(ctx context.Context, in model.ServerCreateInput) (res *v1.ServerCreateRes, err error)
		ServerJoin(ctx context.Context, serverId uint64, link string) (res *v1.ServerJoinRes, err error)
		ServerDel(ctx context.Context, serverId uint64) (res *v1.ServerDelRes, err error)
		ServerModifyName(ctx context.Context, serverId uint64, serverName string) (res *v1.ServerModifyNameRes, err error)
		ServerSearch(ctx context.Context, serverName string) (res *v1.ServerSearchRes, err error)
		ServerInfo(ctx context.Context, serverId uint64) (res *v1.ServerInfoRes, err error)
		ServerInfoUpd(ctx context.Context, in model.ServerInfoUpdInput) (res *v1.ServerInfoUpdRes, err error)
		ServerCount(ctx context.Context, _ *v1.ServerCountReq) (res *v1.ServerCountRes, err error)
		ServerImg(ctx context.Context, serverId uint64, file *ghttp.UploadFile) (res *v1.ServerImgRes, err error)
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
