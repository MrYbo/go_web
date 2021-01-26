package middleware

import (
	"github.com/gin-gonic/gin"
	"login/app/auth"
	"login/app/config"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		conf := config.Conf
		if conf.Server.Authentication == "jwt" {
			auth.AuthJWT(c)
		} else {
			auth.AuthSession(c)
		}
	}
}
