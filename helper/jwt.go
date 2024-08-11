package helper

import (
	"os"
	model "sso-jwt/model/user"
	"sso-jwt/model/web"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user model.User) (string, error) {
	claims := web.SSOClaims{
		IDUser: user.ID,
		Username: user.UserName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "ssojwt.com",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	privateKey, err := os.ReadFile("privateKey")
	PanicIfError(err)

	privateKeyParse, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(privateKey))
	PanicIfError(err)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	tokenString, err := token.SignedString(privateKeyParse)
	PanicIfError(err)

	return tokenString, nil
}