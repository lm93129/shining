package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
	"appdlserver/util/aliyunoss"
)

type DeleteApp struct {
	Id int `form:"id" json:"id" binding:"required" `
}

func (server *DeleteApp) Delete() serializer.Response {
	var appInfo model.AppManage
	err := model.DB.First(&appInfo, server.Id).Error
	if err != nil {
		return serializer.DBErr("不存在该数据", err)
	}

	model.DB.Where("id = ?", server.Id).Delete(&appInfo)
	err = aliyunoss.Bucket.DeleteObject(appInfo.OssUrl)
	if err != nil {
		return serializer.Err(50003, "oss删除失败", err)
	}
	return serializer.Response{
		Code: 200,
		Data: nil,
		Msg:  "删除成功",
	}
}
