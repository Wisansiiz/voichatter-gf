package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta    `path:"/file" method:"post" mime:"multipart/form-data" tags:"file" summary:"file"`
	File      *ghttp.UploadFile `p:"file" v:"required" type:"file" dc:"选择文件"`
	ServerId  int               `p:"serverId" v:"required" dc:"服务器编号"`
	ChannelId int               `p:"channelId" v:"required" dc:"频道编号"`
}

type UploadFileRes struct {
	Url string `json:"url"`
}
