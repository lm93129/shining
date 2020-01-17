package main

import (
	"appdlserver/config"
	"appdlserver/router"
)

func init() {
	config.Init()
}

func main() {

	// 装载路由
	r := router.NewRouter()
	r.Run(":3000")
}
