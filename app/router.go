package app

import (
	"sso-jwt/controller"
	"sso-jwt/exception"
	"sso-jwt/repository"
	"sso-jwt/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"gorm.io/gorm"
)

func NewRouter(DB *gorm.DB, validate *validator.Validate) *httprouter.Router {
	router := httprouter.New()

	router.PanicHandler = exception.ErrorHandler

	NewUserRouter(router, DB, validate)

	return router
}

func NewUserRouter(router *httprouter.Router, DB *gorm.DB, validate *validator.Validate){
	userRepository := repository.NewUserRepository(DB)
	userService := service.NewUserService(userRepository, validate)
	userController := controller.NewUserController(userService)

	router.GET("/api/users", userController.FindAll)
	router.GET("/api/user/:IDUser", userController.FindById)
	router.POST("/api/user/login", userController.Login)
	router.POST("/api/user", userController.Create)
	router.PUT("/api/userpassword/:IDUser", userController.UpdatePassword)
	router.PUT("/api/user/:IDUser", userController.Update)
	router.DELETE("/api/user/:IDUser", userController.Delete)
}