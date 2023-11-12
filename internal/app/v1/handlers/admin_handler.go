package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type AdminHandler interface {
	RegisterHandler(echo.Context) error
}

type AdminHandlerImpl struct {
	AdminService services.AdminService
}

func NewAdminHandler(admin services.AdminService) AdminHandler {
	return &AdminHandlerImpl{
		AdminService: admin,
	}
}

func (handler *AdminHandlerImpl) RegisterHandler(ctx echo.Context) error {
	adminCreateRequest := requests.AdminRequest{}
	err := ctx.Bind(&adminCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.AdminService.RegisterAdmin(ctx, adminCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}
		if strings.Contains(err.Error(), "role admin not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	adminCreateResponse := conversion.AdminDomainToAdminResource(response)

	return responses.StatusCreated(ctx, "Admin created successfully", adminCreateResponse)

}
