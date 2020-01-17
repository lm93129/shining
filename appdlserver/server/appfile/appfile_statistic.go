package appfile

import (
	"appdlserver/model"
	"appdlserver/server/serializer"
)

type Statistic struct {
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
		Group("project_id").
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
