package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"voichatter/internal/model"
)

type UserAvatarReq struct {
	g.Meta `path:"/user/avatar" method:"put" mime:"multipart/form-data" tags:"UserService" summary:""`
	File   *ghttp.UploadFile `p:"file" v:"required" type:"file" dc:"选择头像文件"`
}

type UserAvatarRes struct {
	UserInfo model.UserInfo `json:"userInfo"`
}
