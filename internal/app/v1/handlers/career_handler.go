package handlers

import (
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CareerHandler interface {
	CreateCareer(ctx echo.Context) error
}

type CareerHandlerImpl struct {
	CareerService services.CareerService
}

func NewCareerHandler(career services.CareerService) CareerHandler {
	return &CareerHandlerImpl{
		CareerService: career,
	}
}

func (handler *CareerHandlerImpl) CreateCareer(ctx echo.Context) error {
	carrerCreateRequest := requests.CareerRequest{}
	logo, errLogo := ctx.FormFile("logo")
	cover, errCover := ctx.FormFile("cover")

	if errLogo != nil {
		return exceptions.StatusBadRequest(ctx, errLogo)
	}

	if errCover != nil {
		return exceptions.StatusBadRequest(ctx, errCover)
	}

	err := ctx.Bind(&carrerCreateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.CareerService.CreateCareer(ctx, carrerCreateRequest, logo, cover)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success create career", nil)
}
