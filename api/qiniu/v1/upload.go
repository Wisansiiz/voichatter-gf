package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadFileReq struct {
	g.Meta `path:"/file" method:"post" tags:"file" summary:"file"`
	File   *ghttp.UploadFile `p:"file" v:"required" type:"file" dc:"选择文件"`
}

type UploadFileRes struct {
	Url string `json:"url"`
}
