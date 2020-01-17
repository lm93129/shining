package router

import (
	"appdlserver/router/api"
	"os"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	gin.SetMode(os.Getenv("GIN_MODE"))
	gin.DisableConsoleColor()
	r := gin.New()
	// 中间件, 顺序不能改，从上往下加载
	// 跨域检查配置中间件
	//r.Use(middleware.Cors())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// 健康检查
	r.GET("ping", api.Ping)
	r.MaxMultipartMemory = 300 << 20 // 200MB上传大小限制
	// app相关接口
	AppFile := r.Group("/appFile")
	{
		AppFile.GET("/token", api.UpLoadToOSS)       // 获取oss，token和下载地址
		AppFile.POST("/uploadFile", api.UpLoadFile)  // ci程序上传安装包接口
		AppFile.GET("/appInfo/:id", api.AppInfoShow) // 获取app安装包详情
		AppFile.GET("/appInfos", api.AppFileList)
		AppFile.GET("/project", api.StatisticApp)
		AppFile.GET("/app/list", api.AppFileListOne)
		AppFile.GET("/plist/:id", api.Plist)
		AppFile.DELETE("/deleteFile", api.DeleteFile) // 删除安装包，数据库为软删除
	}
	return r
}
