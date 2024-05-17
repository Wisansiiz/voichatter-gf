package qiniu

import (
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"voichatter/internal/service"

	"voichatter/api/qiniu/v1"
)

func (c *ControllerV1) UploadFile(ctx context.Context, req *v1.UploadFileReq) (res *v1.UploadFileRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, "请上传文件")
	}
	url, err := service.Qiniu().UploadFile(
		ctx,
		req.File,
		gstr.Join([]string{gconv.String(req.ServerId), gconv.String(req.ChannelId)}, "/"),
	)
	if err != nil {
		return nil, err
	}
	return &v1.UploadFileRes{
		Url: url,
	}, nil
}
