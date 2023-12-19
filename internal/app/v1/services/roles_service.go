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

type RoleService interface {
	CreateRole(ctx echo.Context, request requests.RoleRequest) (*domain.Roles, []exceptions.ValidationMessage, error)
	FindRoles(ctx echo.Context) ([]domain.Roles, error)
	DeleteRoleById(ctx echo.Context, id int) error
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

	existingRole, _ := service.RoleRepo.FindByName(request.Name)
	if existingRole != nil {
		return nil, nil, fmt.Errorf("Role already exists")
	}

	role := conversion.RoleCreateRequestToRoleDomain(request)

	result, err := service.RoleRepo.CreateRole(role)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when create role")
	}

	return result, nil, nil
}

func (service *RoleServiceImpl) FindRoles(ctx echo.Context) ([]domain.Roles, error) {
	result, err := service.RoleRepo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("Roles not found")
	}

	return result, nil
}

func (service *RoleServiceImpl) DeleteRoleById(ctx echo.Context, id int) error {
	existingRole, _ := service.RoleRepo.FindById(id)
	if existingRole == nil {
		return fmt.Errorf("Role not found")
	}

	err := service.RoleRepo.DeleteById(id)
	if err != nil {
		return fmt.Errorf("Error when delete role: %s", err.Error())
	}

	return nil

}
