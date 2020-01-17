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

func UpLoadFile(c *gin.Context) {
	var service appfile.UploadFile
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Upload(c)
		c.JSON(200, res)
	}
}

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

func AppInfoShow(c *gin.Context) {
	var service appfile.ShowApp
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

func DeleteFile(c *gin.Context) {
	var service appfile.DeleteApp
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err.Error())
	} else {
		res := service.Delete()
		c.JSON(200, res)
	}
}

func Plist(c *gin.Context) {
	service := appfile.Plist{}
	res := service.Show(c.Param("id"))
	contentLength := len(res)
	extraHeaders := map[string]string{
		"MIME Type": "text/xml",
	}
	c.DataFromReader(200, int64(contentLength), "text/plain; charset=utf-8", bytes.NewReader(res), extraHeaders)
}

func StatisticApp(c *gin.Context) {
	var service appfile.Statistic
	res := service.Statistic()
	c.JSON(200, res)
}
