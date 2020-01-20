package main

import (
	"appdlserver/config"
	"appdlserver/router"
	"fmt"
	"log"
	"os"
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
// @host localhost:3000
func main() {
	// 装载路由
	r := router.InitRouter()

	fmt.Println("|-----------------------------------|")
	fmt.Println("|              闪 灵                |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|     Http Server Start Successful  |")
	fmt.Println("|     Port: " + os.Getenv("PORT") + "  Pid:" + fmt.Sprintf("%d", os.Getpid()) + "         |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	if err := r.Run(":3000"); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
}
