package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
	"appdlserver/util/aliyunoss"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type ShowApp struct {
}

func (server *ShowApp) Show(id string) serializer.Response {
	var appShow model.AppManage
	err := model.DB.First(&appShow, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "不存在",
			Error: err.Error(),
		}
	}
	signedGetURL, err := aliyunoss.Bucket.SignURL(appShow.OssUrl, oss.HTTPGet, 600)
	if err != nil {
		return serializer.Response{
			Code:  50002,
			Msg:   "OSS配置错误",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAppData(appShow,signedGetURL),
	}
}
