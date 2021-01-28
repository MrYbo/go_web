package request

import "github.com/dgrijalva/jwt-go"

type AuthUser struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

type MyClaims struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}
