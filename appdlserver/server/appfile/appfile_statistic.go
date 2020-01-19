package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
)

type Statistic struct {
	StarTime string `form:"star_time" json:"star_time" binding:"required"`
	EndTime  string `form:"end_time" json:"end_time" binding:"required"`
}

type Results struct {
	ProjectId string `json:"project_id"`
	Count     int
}

func (server *Statistic) Statistic() serializer.Response {
	var statistic model.AppManage
	var results []Results
	err := model.DB.Model(&statistic).
		Select("project_id ,COUNT(*)").
		Group("project_id").Where("created_at BETWEEN ? AND ?", server.StarTime, server.EndTime).
		Scan(&results).Error

	if err != nil {
		return serializer.DBErr("", err)
	}

	return serializer.BuildListResponse(BuildStatDatas(results), uint(len(results)))
}

// 序列化返回的json
func BuildStatData(item Results) Results {

	return Results{
		ProjectId: item.ProjectId,
		Count:     item.Count,
	}
}

// 序列化列表
func BuildStatDatas(items []Results) (datas []Results) {
	for _, item := range items {
		data := BuildStatData(item)
		datas = append(datas, data)
	}
	return datas
}
