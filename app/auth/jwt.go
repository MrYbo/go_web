package auth

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	"web_start/app/config"
	"web_start/app/database/mysql"
	"web_start/app/model/entity"
	"web_start/app/schema/request"
	"web_start/app/schema/response"
)

var conf = config.Conf

/**
生成token
*/
func GenToken(id uint, username string) (string, error) {
	claims := request.MyClaims{
		Id:   id,
		Name: username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: conf.Token.Expire * 60 * 60 * time.Now().Unix(),
			Issuer:    conf.Token.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(conf.Token.Secret))

	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

/**
验证token
*/
func ParseToken(tokenStr string) (*request.MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &request.MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Token.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	// 校验token
	if claims, ok := token.Claims.(*request.MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

/**
从请求中获取token
*/
func AuthJWT(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		c.Redirect(http.StatusFound, "/login")
		c.Abort()
		return
	}

	mc, err := ParseToken(token)
	if err != nil {
		response.Failed(c, http.StatusForbidden, "token 验证失败,请重新登录")
		c.Abort()
		return
	}

	var au = request.AuthUser{Id: mc.Id, Name: mc.Name}
	var user entity.User
	mysql.DB.Where("id = ? and name = ?", au.Id, au.Name)
	if user.ID == 0 {
		response.Failed(c, http.StatusBadRequest, "用户不存在")
		c.Abort()
		return
	}
	c.Next()
}

/**
存储token
*/
func SaveAuthJWT(c *gin.Context, user request.AuthUser) string {
	token, err := GenToken(user.Id, user.Name)
	if err != nil {
		response.Failed(c, http.StatusBadRequest, "生成token失败，请重新登录")
		return ""
	}
	return token
}

func HasToken(c *gin.Context) bool {
	token := c.Request.Header.Get("x-token")
	if token == "" {
		return false
	}

	_, err := ParseToken(token)
	if err != nil {
		return false
	}

	return true
}
