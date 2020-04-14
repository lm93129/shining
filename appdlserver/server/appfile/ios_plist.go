package appfile

import (
	"appdlserver/model"
	"appdlserver/util/aliyunoss"
	"appdlserver/util/appfileparser"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type Plist struct {
}

func (server *Plist) Show(id string) (plist []byte, err error){
	var appShow model.AppManage

	err = model.DB.First(&appShow, id).Error
	if err != nil {
		return nil, err
	}
	// 签名获取oss下载地址
	signedGetURL, err := aliyunoss.Bucket.SignURL(appShow.OssUrl, oss.HTTPGet, 600)
	if err!=nil {
		return nil,err
	}
	// 获取plist
	plistdata, _ := appfileparser.EnCodePlist(appShow, signedGetURL)
	return plistdata,nil
}
