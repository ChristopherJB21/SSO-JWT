package service

import (
	"context"
	"sso-jwt/exception"
	"sso-jwt/helper"
	model "sso-jwt/model/user"
	"sso-jwt/repository"

	"github.com/go-playground/validator/v10"
)

type IUserService interface {
	Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse
	Delete(ctx context.Context, IDUser uint)
	Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse
	UpdatePassword(ctx context.Context, request model.UserUpdatePasswordRequest)
	FindAll(ctx context.Context, limit int, offset int) []model.UserResponse
	FindById(ctx context.Context, IDUser uint) model.UserResponse
	Login(ctx context.Context, request model.UserLoginRequest) model.UserLoginResponse
}

type UserService struct {
	UserRepository repository.IUserRepository
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.IUserRepository, validate *validator.Validate) IUserService {
	return &UserService{
		UserRepository: userRepository,
		Validate:       validate,
	}
}

func (service *UserService) Create(ctx context.Context, request model.UserCreateRequest) model.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashPassword, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	newUser := model.User{
		UserName: request.UserName,
		Password: hashPassword,
	}

	newUser = service.UserRepository.Create(ctx, newUser)

	newUser, err = service.UserRepository.FindById(ctx, newUser.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToUserResponse(newUser)
}

func (service *UserService) Delete(ctx context.Context, IDUser uint) {
	user, err := service.UserRepository.FindById(ctx, IDUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.UserRepository.Delete(ctx, user)
}

func (service *UserService) FindAll(ctx context.Context, limit int, offset int) []model.UserResponse {
	users, err := service.UserRepository.FindAll(ctx, limit, offset)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToUsersResponses(users)
}

func (service *UserService) FindById(ctx context.Context, IDUser uint) model.UserResponse {
	findUser, err := service.UserRepository.FindById(ctx, IDUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToUserResponse(findUser)
}

func (service *UserService) Update(ctx context.Context, request model.UserUpdateRequest) model.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	findUser, err := service.UserRepository.FindById(ctx, request.IDUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	findUser.UserName = request.UserName

	findUser = service.UserRepository.Update(ctx, findUser)

	findUser, err = service.UserRepository.FindById(ctx, request.IDUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return model.ToUserResponse(findUser)
}

func (service *UserService) UpdatePassword(ctx context.Context, request model.UserUpdatePasswordRequest) {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	findUser, err := service.UserRepository.FindById(ctx, request.IDUser)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	hashPassword, err := helper.HashPassword(request.NewPassword)
	helper.PanicIfError(err)

	findUser.Password = hashPassword

	findUser = service.UserRepository.Update(ctx, findUser)
}
 
func (service *UserService) Login(ctx context.Context, request model.UserLoginRequest) model.UserLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	findUser, err := service.UserRepository.Login(ctx, request.UserName)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	match := helper.CheckPasswordHash(request.Password, findUser.Password)
	if match {
		token, err := helper.GenerateToken(findUser)
		helper.PanicIfError(err)
		return model.ToUserLoginResponse(findUser, token)
	} else {
		panic(exception.NewAuthenticationError("wrong password"))
	}
}