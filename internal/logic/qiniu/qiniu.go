package qiniu

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/qiniu/go-sdk/v7/auth"
	"github.com/qiniu/go-sdk/v7/storage"
	v1 "voichatter/api/qiniu/v1"
	"voichatter/internal/service"
)

type (
	sQiniu struct{}
)

func init() {
	service.RegisterQiniu(New())
}

func New() service.IQiniu {
	return &sQiniu{}
}

func (s *sQiniu) UploadFile(ctx context.Context, file *ghttp.UploadFile) (req *v1.UploadFileRes, err error) {
	if file == nil {
		return nil, gerror.New("无文件")
	}

	dirPath := "./upload/"
	key, err := file.Save(dirPath, true)
	if err != nil {
		return
	}
	localFile := dirPath + key
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	mac := auth.New(accessKey, secretKey)
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{
		Zone:          &storage.ZoneHuadong,
		UseHTTPS:      true,
		UseCdnDomains: false,
	}
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	err = formUploader.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return
	}
	url := g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	g.Dump(url)
	if err = gfile.Remove(localFile); err != nil {
		return nil, gerror.New("删除文件失败")
	}
	return &v1.UploadFileRes{
		Url: url,
	}, nil
}
