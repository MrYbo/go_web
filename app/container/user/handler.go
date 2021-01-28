package user

import (
	"encoding/gob"
	"github.com/gin-gonic/gin"
	"net/http"
	"web_start/app/auth"
	"web_start/app/config"
	"web_start/app/database/mysql"
	"web_start/app/model/entity"
	"web_start/app/schema/request"
	"web_start/app/schema/response"
)

var conf = config.Conf

func LoginPage(c *gin.Context) {
	// 鉴权
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	var lp request.LoginParams
	gob.Register(request.AuthUser{})
	if config.Conf.Server.Authentication == "jwt" {
		if auth.HasToken(c) {
			c.Redirect(http.StatusFound, config.Conf.Redirect.Url)
			return
		}
	} else {
		if auth.HasSession(c) {
			c.Redirect(http.StatusFound, config.Conf.Redirect.Url)
			return
		}
	}

	if err := request.Validate(c, &lp); err != nil {
		return
	}

	var user entity.User
	mysql.DB.Where("name=?", lp.Username)

	if user.ID == 0 || user.Password != lp.Password {
		response.Failed(c, http.StatusBadRequest, "账号或密码错误")
		return
	}

	authUser := request.AuthUser{Id: user.ID, Name: user.Name}

	if config.Conf.Server.Authentication == "jwt" {
		// 登录成功，生成token凭证
		token := auth.SaveAuthJWT(c, authUser)
		response.Success(c, http.StatusOK, map[string]interface{}{
			"token":    token,
			"redirect": conf.Redirect.Url,
		})
		return
	}

	auth.SaveAuthSession(c, authUser)
	response.Success(c, http.StatusOK, map[string]interface{}{
		"user": user,
	})
}
