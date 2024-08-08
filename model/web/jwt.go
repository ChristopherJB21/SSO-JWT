package web

import "github.com/golang-jwt/jwt/v5"

type SSOClaims struct {
	IDUser   uint   `json:"iduser"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}
