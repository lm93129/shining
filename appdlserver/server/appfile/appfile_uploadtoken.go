package appfile

import (
	"appdlserver/server/serializer"
	"appdlserver/util/aliyunoss"
	"fmt"
	"mime"
	"path/filepath"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type UploadTokenService struct {
	Filename string `form:"filename" json:"filename"`
}

// Post 创建token
func (service *UploadTokenService) Post() serializer.Response {

	// 获取扩展名
	var appext string
	ext := filepath.Ext(service.Filename)
	if ext == ".apk" {
		appext = "android"
	} else {
		appext = "ios"
	}
	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	key := fmt.Sprintf("upload/appfile/%s/%s%s",
		appext,
		time.Now().Format("2006-01-02 15:04:05"),
		ext,
	)
	// 签名直传。
	signedPutURL, err := aliyunoss.Bucket.SignURL(key, oss.HTTPPut, 600, options...)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}
	// 查看
	signedGetURL, err := aliyunoss.Bucket.SignURL(key, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: map[string]string{
			"key": key,
			"put": signedPutURL,
			"get": signedGetURL,
		},
	}
}
