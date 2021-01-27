package user

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	group := router.Group("")
	{
		group.GET("/login", LoginPage)
		group.POST("/login", Login)
	}
}
