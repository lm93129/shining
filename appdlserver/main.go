package main

import (
	"appdlserver/config"
	"appdlserver/router"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
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
		Addr:           fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   15 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("|-----------------------------------|")
	fmt.Println("|              闪 灵                |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("|     Http Server Start Successful  |")
	fmt.Println("|     Port: " + os.Getenv("PORT") + "  Pid:" + fmt.Sprintf("%d", os.Getpid()) + "         |")
	fmt.Println("|-----------------------------------|")
	fmt.Println("")

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server listen: %s\n", err)
		}
	}()

	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, os.Interrupt)
	sig := <-signalChan
	log.Println("Get Signal:", sig)
	log.Println("Shutdown Server ...")
	// 接收到关闭信号后等待3秒退出
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
