package config

import (
	"appdlserver/model"
	"appdlserver/util"
	"appdlserver/util/aliyunoss"
	"os"

	"github.com/joho/godotenv"
)

// Init 初始化配置项
func Init() {
	// 从本地读取环境变量
	_ = godotenv.Load()

	// 设置日志级别
	util.BuildLogger(os.Getenv("LOG_LEVEL"))

	// 连接数据库
	model.Database()
	// 连接Redis缓存
	//model.Redis()
	aliyunoss.InitOss()
}
