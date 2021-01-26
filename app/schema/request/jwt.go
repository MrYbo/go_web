package request

import "github.com/dgrijalva/jwt-go"

type AuthUser struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

type MyClaims struct {
	Id int `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

