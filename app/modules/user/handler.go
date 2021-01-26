package user

import (
	"encoding/gob"
	"login/app/auth"
	"login/app/config"
	"login/app/model/dao"
	"login/app/model/entity"
	"login/app/schema/request"
	"login/app/schema/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var userDAO *dao.User
var conf = config.Conf

func LoginPage(c *gin.Context) {
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

	user := userDAO.FindOne("name=?", lp.Username).(entity.User)

	if user.Id == 0 || user.Password != lp.Password {
		response.Failed(c, http.StatusBadRequest, "账号或密码错误")
		return
	}

	authUser := request.AuthUser{Id: user.Id, Name: user.Name}

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
		"redirect": conf.Redirect.Url,
	})
}
