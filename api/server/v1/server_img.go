package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"voichatter/internal/model"
)

type ServerImgReq struct {
	g.Meta   `path:"/server/img/{serverId}" method:"put" mime:"multipart/form-data" tags:"ServerService" summary:""`
	ServerId uint64            `p:"serverId" v:"required"`
	File     *ghttp.UploadFile `p:"file"     v:"required" type:"file" dc:"选择头像文件"`
}

type ServerImgRes struct {
	ServerInfo *model.Server `json:"serverInfo"`
}
