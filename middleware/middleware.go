package middleware

import (
	"net/http"
	"sso-jwt/helper"
	"sso-jwt/model/web"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Middleware struct {
	Handler  http.Handler
	Validate *validator.Validate
}

func NewMiddleware(handler http.Handler, validate *validator.Validate) *Middleware {
	return &Middleware{
		Handler:  handler,
		Validate: validate,
	}
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	DecodedAPIKey := GetAppKey(request)

	if DecodedAPIKey != viper.GetString("apiKey") {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		helper.WriteToResponseBody(writer, webResponse)

		return
	}

	middleware.Handler.ServeHTTP(writer, request)
}

func GetAppKey(request *http.Request) string {
	var APIKey = request.Header.Get("X-API-Key")

	return APIKey
}
