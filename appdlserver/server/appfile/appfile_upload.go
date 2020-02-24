package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
	"appdlserver/util"
	"appdlserver/util/aliyunoss"
	"appdlserver/util/appfileparser"
	"mime"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/gin-gonic/gin"
)

type UploadFile struct {
	UpdateDescription string `json:"update_description" from:"UpdateDescription"`
	ProjectId         string `json:"project_id" from:"ProjectId" binding:"required"`
}

// Post 创建token
func (service *UploadFile) Upload(c *gin.Context) serializer.Response {
	file, fileh, _ := c.Request.FormFile("upload")

	var appExt string
	var dst string
	var appFileInfo *appfileparser.AppInfo

	// 获取扩展名
	ext := filepath.Ext(fileh.Filename)
	if ext == ".apk" {
		appExt = "android"
	} else {
		appExt = "ios"
	}

	dst = "/tmp/" + util.RandStringRunes(4) + fileh.Filename
	// win调试用
	//dst = util.RandStringRunes(4) + fileh.Filename

	err := c.SaveUploadedFile(fileh, dst)

	if err != nil {
		return serializer.ParamErr("安装包缓存失败", err)
	}
	appFileInfo, _ = appfileparser.NewAppParser(dst)
	defer os.Remove(dst)

	options := []oss.Option{
		oss.ContentType(mime.TypeByExtension(ext)),
	}

	fileName := strconv.FormatInt(time.Now().Unix(), 10)
	key := "upload/appfile/" + service.ProjectId + "/" + appExt + "/" + fileName

	appdb := model.AppManage{
		Name:              appFileInfo.Name,
		BundleId:          appFileInfo.BundleId,
		Version:           appFileInfo.Version,
		Build:             appFileInfo.Build,
		Size:              appFileInfo.Size,
		OssUrl:            key + ext,
		AppType:           appExt,
		ProjectId:         service.ProjectId,
		UpdateDescription: service.UpdateDescription,
	}

	// 上传到oss
	err = aliyunoss.Bucket.PutObject(key+ext, file, options...)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "上传失败",
			Error: err.Error(),
		}
	}
	// 写入数据库记录
	if err := model.DB.Create(&appdb).Error; err != nil {
		return serializer.ParamErr("写入失败", err)
	}

	return serializer.Response{
		Code: 200,
		Data: key,
		Msg:  "上传成功",
	}
}
