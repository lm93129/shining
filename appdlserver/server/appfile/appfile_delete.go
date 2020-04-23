package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
	"appdlserver/util/aliyunoss"
)

type DeleteApp struct {
	// 要删除的安装包ID
	Id []int `form:"id" json:"id" binding:"required" `
}

func (server *DeleteApp) Delete() serializer.Response {
	var appInfo []model.AppManage
	model.DB.Where(server.Id).Find(&appInfo)
	// 异步删除oss中的数据，暂不处理oss返回的错误
	go func() {
		for _, v := range appInfo {
			_ = aliyunoss.Bucket.DeleteObject(v.OssUrl)
		}
	}()

	err := model.DB.Where(server.Id).Delete(model.AppManage{}).Error
	if err != nil {
		return serializer.DBErr("数据删除失败", err)
	}

	return serializer.Response{
		Code: 200,
		Data: nil,
		Msg:  "删除成功",
	}
}
