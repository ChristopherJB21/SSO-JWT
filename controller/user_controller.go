package controller

import (
	"encoding/json"
	"net/http"
	"sso-jwt/helper"
	model "sso-jwt/model/user"
	"sso-jwt/model/web"
	"sso-jwt/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type IUserController interface {
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}

type UserController struct {
	UserService service.IUserService
}

func NewUserController(userService service.IUserService) IUserController {
	return &UserController{
		UserService: userService,
	}
}

func (controller *UserController) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userCreateRequest := model.UserCreateRequest{}
	var result interface{} = &userCreateRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(result)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Create(request.Context(), userCreateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDUser := params.ByName("IDUser")
	id, err := strconv.ParseUint(IDUser, 10, 64)
	helper.PanicIfError(err)

	controller.UserService.Delete(request.Context(), uint(id))
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	var limit, offset int

	limitQuery, err := helper.ReadFromQueryParams("limit", request)
	if err != nil {
		limit = 10
	} else {
		limit, err = strconv.Atoi(limitQuery)
		helper.PanicIfError(err)
	}

	offsetQuery, err := helper.ReadFromQueryParams("offset", request)
	if err != nil {
		offset = 0
	} else {
		offset, err = strconv.Atoi(offsetQuery)
		helper.PanicIfError(err)
	}

	userResponses := controller.UserService.FindAll(request.Context(), limit, offset)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   userResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	IDUser := params.ByName("IDUser")
	id, err := strconv.ParseUint(IDUser, 10, 64)
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(request.Context(), uint(id))

	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := model.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	IDUser := params.ByName("IDUser")
	id, err := strconv.ParseUint(IDUser, 10, 64)
	helper.PanicIfError(err)

	userUpdateRequest.IDUser = uint(id)

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   userResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) Login(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userLoginRequest := model.UserLoginRequest{}
	helper.ReadFromRequestBody(request, &userLoginRequest)

	userLoginResponse := controller.UserService.Login(request.Context(), userLoginRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
		Data:   userLoginResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserController) UpdatePassword(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdatePasswordRequest := model.UserUpdatePasswordRequest{}
	helper.ReadFromRequestBody(request, &userUpdatePasswordRequest)

	IDUser := params.ByName("IDUser")
	id, err := strconv.ParseUint(IDUser, 10, 64)
	helper.PanicIfError(err)

	userUpdatePasswordRequest.IDUser = uint(id)

	controller.UserService.UpdatePassword(request.Context(), userUpdatePasswordRequest)
	webResponse := web.WebResponse{
		Code:   http.StatusAccepted,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}