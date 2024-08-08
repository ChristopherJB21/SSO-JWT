package main

import (
	"log"
	"net/http"
	"sso-jwt/app"
	"sso-jwt/helper"
	"sso-jwt/middleware"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

func main() {
	app.NewViper()

	DB := app.NewDB()

	validate := validator.New()
	
	router := app.NewRouter(DB, validate)

	server := http.Server{
		Addr:    viper.GetString("server.addr"),
		Handler: middleware.NewMiddleware(router, validate),
	}

	log.Println(viper.GetString("appName") + " Application Start")
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
