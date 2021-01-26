package router

import (
	"github.com/gin-gonic/gin"
	"login/app/config"
	"login/app/middleware"
	"login/app/modules/user"
	"net/http"
)

type routerGroup func(engine *gin.Engine)

var routers = [...]routerGroup{
	user.Router,
}



func Init(engine *gin.Engine) *gin.Engine {
	// 配置模板
	engine.LoadHTMLGlob("app/view/*")

	engine.NoRoute(middleware.Auth(), func(c *gin.Context) {
		// 登录成功之后访问任何页面都跳转
		c.Redirect(http.StatusFound, config.Conf.Redirect.Url)
		return
	})

	// 路由初始化
	for _, router := range routers {
		router(engine)
	}
	return engine
}
