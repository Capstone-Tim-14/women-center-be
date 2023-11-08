package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/conversion"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, error)
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

func (service *UserServiceImpl) RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingUser, _ := service.UserRepo.FindyByEmail(request.Email)
	if existingUser != nil {
		return nil, fmt.Errorf("email already exist")
	}

	user := conversion.UserCreateRequestToUserDomain(request)

	user.Password = helpers.HashPassword(request.Password)

	result, err := service.UserRepo.CreateUser(user)
	if err != nil {
		return nil, fmt.Errorf("Error when register user: %s", err.Error())
	}

	return result, nil
}
