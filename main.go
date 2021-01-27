package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"web_start/app/config"
	"web_start/app/database/mysql"
	"web_start/app/database/redis"
	"web_start/app/middleware"
	"web_start/app/router"
)

var app = gin.New()
var conf = config.Conf

func init() {
	gin.SetMode(gin.ReleaseMode)
	mysql.Init()
	redis.Init()
	middleware.Init(app)
	router.Init(app)
}

func main() {
	RunServer()
}

func RunServer() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: app,
	}

	log.Println(fmt.Sprintf("Listening and serving HTTP on Port: %d, Pid: %d", config.Conf.Server.Port, os.Getpid()))

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server listen: %s\n", err)
		}
	}()

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// 创建5s的超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exiting")
}
