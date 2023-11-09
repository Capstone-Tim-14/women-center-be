package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/conversion"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	CreateRoleHandler(echo.Context) error
}

type RoleHandlerImpl struct {
	RoleService services.RoleService
}

func NewRoleHandler(role services.RoleService) RoleHandler {
	return &RoleHandlerImpl{
		RoleService: role,
	}
}

func (handler *RoleHandlerImpl) CreateRoleHandler(ctx echo.Context) error {
	roleCreateRequest := requests.RoleRequest{}
	err := ctx.Bind(&roleCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, err := handler.RoleService.CreateRole(ctx, roleCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}
	roleCreateResponse := conversion.RoleDomainToRoleResponse(response)
	return responses.StatusCreated(ctx, "Role created successfully", roleCreateResponse)

}
