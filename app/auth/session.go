package auth

import (
	"login/app/config"
	"login/app/schema/request"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init() sessions.Store {
	conf := config.Conf
	var Store sessions.Store
	if conf.Storage == "redis" {
		Store, _ = redis.NewStore(
			conf.Size,
			"tcp",
			conf.Addr,
			conf.Redis.Password,
			[]byte(conf.Session.Key))
	} else {
		Store = cookie.NewStore([]byte(conf.Session.Key))
		Store.Options(sessions.Options{
			Path:     conf.Session.Path,
			Domain:   conf.Session.Domain,
			MaxAge:   conf.Session.MaxAge,
			HttpOnly: conf.Session.HttpOnly,
		})
	}
	return Store
}

//保存session信息
func SaveAuthSession(c *gin.Context, user request.AuthUser) {
	session := sessions.Default(c)
	session.Set("user", user)
	err := session.Save()
	if err != nil {
		logrus.Error("session save error", err)
	}
}

// 退出时清除session
func ClearAuthSession(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	_ = session.Save()
}

func HasSession(c *gin.Context) bool {
	session := sessions.Default(c)
	if sessionValue := session.Get("user"); sessionValue == nil {
		return false
	}
	return true
}

func AuthSession(c *gin.Context) {
	session := sessions.Default(c)
	sessionValue := session.Get("user")
	if sessionValue == nil {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}
}
