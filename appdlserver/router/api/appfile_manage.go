package api

import (
	"appdlserver/server/appfile"
	"bytes"

	"github.com/gin-gonic/gin"
)

func UpLoadToOSS(c *gin.Context) {
	var service appfile.UploadTokenService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Post()
		c.JSON(200, res)
	}
}

// @Summary 上传安装包文件
// @Description 上传安装包文件
// @Produce json
// @tags APP安装包相关
// @Param ProjectId query string true "安装包所属项目ID"
// @Param UpdateDescription query string true "安装包更新说明"
// @Param upload formData file true "安装包文件"
// @Success 200 {object} serializer.Response
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/uploadFile [post]
func UpLoadFile(c *gin.Context) {
	var service appfile.UploadFile
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Upload(c)
		c.JSON(200, res)
	}
}

// @Summary 获取安装包列表
// @Description 获取安装包列表
// @Produce json
// @tags APP安装包相关
// @Param body body appfile.ListApp true "安装包列表"
// @Success 200 {object} serializer.Response
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/appInfos [get]
func AppFileList(c *gin.Context) {
	var service appfile.ListApp
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.List()
		c.JSON(200, res)
	}
}

func AppFileListOne(c *gin.Context) {
	var service appfile.ListApp
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.ListOne()
		c.JSON(200, res)
	}
}

// @Summary 获取单个安装包详情
// @Description 获取单个安装包详情
// @Produce json
// @tags APP安装包相关
// @Param id path string true "安装包ID"
// @Success 200 {object} serializer.Response
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/appInfo/{id} [get]
func AppInfoShow(c *gin.Context) {
	var service appfile.ShowApp
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

// @Summary 删除安装包
// @Description 删除安装包
// @Produce json
// @tags APP安装包相关
// @Param id query string true "安装包id"
// @Success 200 {object} serializer.Response
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/deleteFile [delete]
func DeleteFile(c *gin.Context) {
	var service appfile.DeleteApp
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Delete()
		c.JSON(200, res)
	}
}

// @Summary IOS获取plist文件
// @Description IOS获取plist文件
// @Produce xml
// @tags APP安装包相关
// @Param id path string true "安装包ID"
// @Success 200 {object} appfileparser.sparseBundleHeader
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/plist/{id} [get]
func Plist(c *gin.Context) {
	service := appfile.Plist{}
	res := service.Show(c.Param("id"))
	contentLength := len(res)
	extraHeaders := map[string]string{
		"MIME Type": "text/xml",
	}
	c.DataFromReader(200, int64(contentLength), "text/plain; charset=utf-8", bytes.NewReader(res), extraHeaders)
}

// @Summary APP安装包统计
// @Description APP安装包统计
// @Produce json
// @tags APP安装包相关
// @Param star_time query string true "起始时间" default 2020-01-02
// @Param end_time query string true "结束事件" default 2020-02-02
// @Success 200 {object} serializer.Response
// @Failure 500 {string} string "err_code：50001 数据库操作失败 err_code: 40001 参数错误"
// @Router /appFile/project [get]
func StatisticApp(c *gin.Context) {
	var service appfile.Statistic
	if err := c.ShouldBind(&service); err == nil {
		res := service.Statistic()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}
