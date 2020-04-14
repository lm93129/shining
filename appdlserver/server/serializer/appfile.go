package serializer

import (
	"appdlserver/model"
	"appdlserver/util/aliyunoss"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

type AppData struct {
	ID                uint   `json:"id"`
	CreatedTime       string `json:"created_time"`
	AppName           string `json:"app_name"`
	AppBundleId       string `json:"app_bundle_id"`
	AppVersion        string `json:"app_version"`
	AppBuild          string `json:"app_build"`
	AppSize           int64  `json:"app_size"`
	OssUrl            string `json:"oss_url"`
	VersionType       string `json:"version_type"`
	ProjectId         string `json:"project_id"`
	UpdateDescription string `json:"update_description"`
}

// 序列化返回的json
func BuildAppData(item model.AppManage,downUrl string) AppData {
	if item.AppType == "ios" {
		downUrl = fmt.Sprintf("itms-services://?action=download-manifest&url=%s/appFile/plist/%d",
			os.Getenv("BASE_URL"),
			item.ID,
		)
	}
	return AppData{
		ID:                item.ID,
		CreatedTime:       item.CreatedAt.In(cstZone).Format("2006-01-02 15:04:05"),
		AppName:           item.Name,
		AppBundleId:       item.BundleId,
		AppVersion:        item.Version,
		AppBuild:          item.Build,
		AppSize:           item.Size >> 20,
		OssUrl:            downUrl,
		VersionType:       item.AppType,
		ProjectId:         item.ProjectId,
		UpdateDescription: item.UpdateDescription,
	}
}

func BuildAppNoUrlData(item model.AppManage) AppData {
	return AppData{
		ID:                item.ID,
		CreatedTime:       item.CreatedAt.In(cstZone).Format("2006-01-02 15:04:05"),
		AppName:           item.Name,
		AppBundleId:       item.BundleId,
		AppVersion:        item.Version,
		AppBuild:          item.Build,
		AppSize:           item.Size >> 20,
		VersionType:       item.AppType,
		ProjectId:         item.ProjectId,
		UpdateDescription: item.UpdateDescription,
	}
}

// 序列化列表
func BuildAppDatas(items []model.AppManage) (appdatas []AppData) {
	for _, item := range items {
		signedGetURL, _ := aliyunoss.Bucket.SignURL(item.OssUrl, oss.HTTPGet, 600)

		appdata := BuildAppData(item,signedGetURL)
		appdatas = append(appdatas, appdata)
	}
	return appdatas
}

func BuildAppNoUrlDatas(items []model.AppManage) (appdatas []AppData) {
	for _, item := range items {
		appdata := BuildAppNoUrlData(item)
		appdatas = append(appdatas, appdata)
	}
	return appdatas
}
