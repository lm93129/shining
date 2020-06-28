package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
)

type ListApp struct {
	PageSize    int    `form:"page_size"`
	Page        int    `form:"page"`
	ProjectId   string `form:"project_id"`
	VersionType string `form:"version_type"`
}

func (server *ListApp) List() serializer.Response {
	var list []model.AppManage
	type Result struct {
		ProjectId string
	}
	var result []Result
	var index = 0
	total := 0
	if server.PageSize == 0 {
		server.PageSize = 10
	}
	if server.Page > 0 {
		index = (server.Page - 1) * server.PageSize
	}

	if server.ProjectId == "" {
		if err := model.DB.Table("app_manage").
			Select("project_id").
			Group("project_id").
			Scan(&result).Error; err != nil {
			return serializer.DBErr("", err)
		}

		return serializer.Response{
			Code: 200,
			Data: result,
		}
	}

	if server.VersionType == "" {
		if err := model.DB.Order("id desc").Where("project_id = ?", server.ProjectId).Limit(server.PageSize).Offset(index).Find(&list).Error; err != nil {
			return serializer.DBErr("", err)
		}
		if err := model.DB.Where("project_id = ?", server.ProjectId).Model(model.AppManage{}).Count(&total).Error; err != nil {
			return serializer.DBErr("", err)
		}
	} else {
		if err := model.DB.Order("id desc").Where("project_id = ? AND app_type = ?", server.ProjectId, server.VersionType).Limit(server.PageSize).Offset(index).Find(&list).Error; err != nil {
			return serializer.DBErr("", err)
		}
		if err := model.DB.Where("project_id = ? AND app_type = ?", server.ProjectId, server.VersionType).Model(model.AppManage{}).Count(&total).Error; err != nil {
			return serializer.DBErr("", err)
		}
	}

	return serializer.BuildListResponse(serializer.BuildAppDatas(list), uint(total))
}

func (server *ListApp) ListOne() serializer.Response {
	var list []model.AppManage
	var index = 0
	total := 0
	if server.PageSize == 0 {
		server.PageSize = 10
	}
	if server.Page > 0 {
		index = (server.Page - 1) * server.PageSize
	}

	if err := model.DB.Where("project_id = ? AND app_type = ?", server.ProjectId, server.VersionType).Model(model.AppManage{}).Count(&total).Error; err != nil {
		serializer.DBErr("", err)
	}
	if err := model.DB.Order("id desc").Where("project_id = ? AND app_type = ?", server.ProjectId, server.VersionType).Limit(server.PageSize).Offset(index).Find(&list).Error; err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.BuildListResponse(serializer.BuildAppNoUrlDatas(list), uint(total))
}
