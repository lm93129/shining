package main

import (
	"appdlserver/config"
	"appdlserver/router"
	"appdlserver/util"
	"fmt"
	"net/http"
	"os"
	"time"
)

func init() {
	config.Init()
}

// @title shining安装包管理系统
// @version 1.0
// @description shining一个APP安装包管理系统

// @contact.name lm93129
// @contact.url https://fs.tn
// @contact.email lm93129@qq.com

//@host localhost:3000
func main() {
	// 装载路由
	r := router.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%s", os.Getenv("PROT")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 2 << 20,
	}

	util.Log().Println("DataCollectApi服务启动\n")
	if err := s.ListenAndServe(); err != nil {
		util.Log().Println("DataCollectApi服务关闭\n")
	}
}
