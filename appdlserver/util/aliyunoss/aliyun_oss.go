package aliyunoss

import (
	"appdlserver/util"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var Bucket *oss.Bucket

func InitOss() {
	client, err := oss.New(os.Getenv("OSS_END_POINT"), os.Getenv("OSS_ACCESS_KEY_ID"), os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		util.Log().Panic("OSS client初始化失败", err)
	}

	// 获取存储空间。
	bucket, err := client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		util.Log().Panic("OSS bucket初始化失败", err)
	}
	Bucket = bucket
}
