package services

import (
	"woman-center-be/internal/app/repositories"

	"github.com/go-playground/validator/v10"
)

type UserService interface {
	RegisterUser()
}

type UserServiceImpl struct {
	UserRepo  repositories.UserRepository
	validator *validator.Validate
}

func NewUserService(user repositories.UserRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo:  user,
		validator: validator,
	}
}

func (user *UserServiceImpl) RegisterUser() {

}
