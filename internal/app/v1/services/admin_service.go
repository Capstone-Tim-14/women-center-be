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

type AdminService interface {
	RegisterAdmin(ctx echo.Context, request requests.AdminRequest) (*domain.Admin, []exceptions.ValidationMessage, error)
}

type AdminServiceImpl struct {
	AdminRepo repositories.AdminRepository
	RoleRepo  repositories.RoleRepository
	validator *validator.Validate
}

func NewAdminService(admin repositories.AdminRepository, validator *validator.Validate, role repositories.RoleRepository) AdminService {
	return &AdminServiceImpl{
		AdminRepo: admin,
		RoleRepo:  role,
		validator: validator,
	}
}

func (service *AdminServiceImpl) RegisterAdmin(ctx echo.Context, request requests.AdminRequest) (*domain.Admin, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingAdmin, _ := service.AdminRepo.FindyByEmail(request.Email)
	if existingAdmin != nil {
		return nil, nil, fmt.Errorf("Email already exists")
	}

	getRoleAdmin, _ := service.RoleRepo.FindByName("admin")
	if getRoleAdmin == nil {
		return nil, nil, fmt.Errorf("role admin not found")
	}

	request.Role_id = uint(getRoleAdmin.Id)

	admin := conversion.AdminCreateRequestToAdminDomain(request)

	admin.Credential.Password = helpers.HashPassword(request.Password)

	result, err := service.AdminRepo.CreateAdmin(admin)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when register admin: %s", err.Error())
	}

	return result, nil, nil
}
