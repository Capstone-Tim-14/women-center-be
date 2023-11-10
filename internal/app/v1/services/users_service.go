package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error)
}

type UserServiceImpl struct {
	UserRepo  repositories.UserRepository
	RoleRepo  repositories.RoleRepository
	validator *validator.Validate
}

func NewUserService(user repositories.UserRepository, role repositories.RoleRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo:  user,
		RoleRepo:  role,
		validator: validator,
	}
}

func (service *UserServiceImpl) RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingUser, _ := service.UserRepo.FindyByEmail(request.Email)
	if existingUser != nil {
		return nil, nil, fmt.Errorf("email already exist")
	}

	getRoleUser, _ := service.RoleRepo.FindByName("user")
	if getRoleUser == nil {
		return nil, nil, fmt.Errorf("role user not found")
	}

	request.Role_id = uint(getRoleUser.Id)

	user := conversion.UserCreateRequestToUserDomain(request)

	user.Credential.Password = helpers.HashPassword(request.Password)

	result, err := service.UserRepo.CreateUser(user)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when register user: %s", err.Error())
	}

	return result, nil, nil
}
