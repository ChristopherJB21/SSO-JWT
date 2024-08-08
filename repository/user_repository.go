package repository

import (
	"context"
	"errors"
	"sso-jwt/exception"
	"sso-jwt/helper"
	model "sso-jwt/model/user"
	"strings"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, user model.User) model.User
	Delete(ctx context.Context, user model.User)
	Update(ctx context.Context, user model.User) model.User
	FindAll(ctx context.Context, limit int, offset int) ([]model.User, error)
	FindById(ctx context.Context, IDUser uint) (model.User, error)
	Login(ctx context.Context, Username string) (model.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (repository *UserRepository) Create(ctx context.Context, user model.User) model.User {
	result := repository.DB.Create(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key"){
		panic(exception.NewBadRequestError("username already used"))
	}

	helper.PanicIfError(result.Error)

	return user
}

func (repository *UserRepository) Delete(ctx context.Context, user model.User) {
	result := repository.DB.Delete(&user)

	helper.PanicIfError(result.Error)
}

func (repository *UserRepository) FindAll(ctx context.Context, limit int, offset int) ([]model.User, error) {
	var users []model.User

	result := repository.DB.Limit(limit).Offset(offset).Find(&users)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) || len(users) < 1 {
		return users, errors.New("users are not found")
	}

	helper.PanicIfError(result.Error)

	return users, nil
}

func (repository *UserRepository) FindById(ctx context.Context, IDUser uint) (model.User, error) {
	var user model.User

	result := repository.DB.First(&user, IDUser)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("user is not found")
	}

	helper.PanicIfError(result.Error)

	return user, nil
}

func (repository *UserRepository) Update(ctx context.Context, user model.User) model.User {
	result := repository.DB.Save(&user)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key"){
		panic(exception.NewBadRequestError("username already used"))
	}

	helper.PanicIfError(result.Error)

	return user
}

func (repository *UserRepository) Login(ctx context.Context, Username string) (model.User, error) {
	var user model.User

	result := repository.DB.Where(&model.User{UserName: Username}).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return user, errors.New("username is not found")
	}

	helper.PanicIfError(result.Error)

	return user, nil
}