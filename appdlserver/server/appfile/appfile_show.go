package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
)

type ShowApp struct {
}

func (server *ShowApp) Show(id string) serializer.Response {
	var appshow model.AppManage
	err := model.DB.First(&appshow, id).Error
	if err != nil {
		return serializer.Response{
			Code:  404,
			Msg:   "不存在",
			Error: err.Error(),
		}
	}
	return serializer.Response{
		Code: 200,
		Data: serializer.BuildAppData(appshow),
	}
}
