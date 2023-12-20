package handlers

import (
	"strconv"
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type RoleHandler interface {
	CreateRoleHandler(echo.Context) error
	FindRolesHandler(echo.Context) error
	DeleteRoleByIdHandler(echo.Context) error
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

	response, validation, err := handler.RoleService.CreateRole(ctx, roleCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Role already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}
	roleCreateResponse := conversion.RoleDomainToRoleResource(response)

	return responses.StatusCreated(ctx, "Role created successfully", roleCreateResponse)
}

func (handler *RoleHandlerImpl) FindRolesHandler(ctx echo.Context) error {
	response, err := handler.RoleService.FindRoles(ctx)
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	roleResponse := conversion.ConvertRoleResource(response)

	return responses.StatusOK(ctx, "Success Get All Role", roleResponse)

}

func (handler *RoleHandlerImpl) DeleteRoleByIdHandler(ctx echo.Context) error {
	roleId := ctx.Param("id")
	roleIdInt, err := strconv.Atoi(roleId)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	err = handler.RoleService.DeleteRoleById(ctx, roleIdInt)
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success Delete Role By Id", nil)

}
