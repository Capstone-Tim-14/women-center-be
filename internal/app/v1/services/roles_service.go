package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/web/conversion"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RoleService interface {
	CreateRole(ctx echo.Context, request requests.RoleRequest) (*domain.Roles, []exceptions.ValidationMessage, error)
}

type RoleServiceImpl struct {
	RoleRepo  repositories.RoleRepository
	validator *validator.Validate
}

func NewRoleService(role repositories.RoleRepository, validator *validator.Validate) RoleService {
	return &RoleServiceImpl{
		RoleRepo:  role,
		validator: validator,
	}
}

func (service *RoleServiceImpl) CreateRole(ctx echo.Context, request requests.RoleRequest) (*domain.Roles, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	role := conversion.RoleCreateRequestToRoleDomain(request)

	result, err := service.RoleRepo.CreateRole(role)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when create role: %s", err.Error())
	}

	return result, nil, nil
}
