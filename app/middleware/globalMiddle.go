package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"login/app/auth"
	"login/app/config"
)

/**
	全局中间件初始化
 */

func Init(router *gin.Engine) {
	conf := config.Conf
	router.Use(
		gin.Recovery(),
		cors.Default(),
		LoggerToFile(),
		)

	if conf.Server.Authentication == "session"{
		router.Use(sessions.Sessions("session_id", auth.Init()))
	}
}


