package api

import (
	"appdlserver/server/serializer"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary 健康检查
// @tags 健康检查
// @Description 健康检查接口，如果返回200则表示成功
// @Accept json
// @Produce json
// @Success 200 {string} string "pong"
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {

	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.ParamErr("JSON类型不匹配", err)
	}

	return serializer.ParamErr("参数错误", err)
}
