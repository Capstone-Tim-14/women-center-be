package handlers

import (
	"strconv"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"

	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CareerHandler interface {
	CreateCareer(ctx echo.Context) error
	FindAllCareer(ctx echo.Context) error
	FindDetailCareer(ctx echo.Context) error
	UpdateCareer(ctx echo.Context) error
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

func (handler *CareerHandlerImpl) FindAllCareer(ctx echo.Context) error {
	career, err := handler.CareerService.FindAllCareer(ctx)

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	careerResource := conversion.ConvertCareerRsource(career)

	return responses.StatusOK(ctx, "Success get career", careerResource)
}

func (handler *CareerHandlerImpl) FindDetailCareer(ctx echo.Context) error {

	getId := ctx.Param("id")
	detailId, err := strconv.Atoi(getId)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	career, err := handler.CareerService.FindCareerByid(ctx, detailId)
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	careerResource := conversion.ConvertCareerDetailResource(career)

	return responses.StatusOK(ctx, "Success get career", careerResource)
}

func (handler *CareerHandlerImpl) UpdateCareer(ctx echo.Context) error {
	careerUpdateRequest := requests.CareerRequest{}
	errBinding := ctx.Bind(&careerUpdateRequest)
	logo, errLogo := ctx.FormFile("logo")
	cover, errCover := ctx.FormFile("cover")

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	if errLogo != nil {
		return exceptions.StatusBadRequest(ctx, errLogo)
	}

	if errCover != nil {
		return exceptions.StatusBadRequest(ctx, errCover)
	}

	err := ctx.Bind(&careerUpdateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	validation, err := handler.CareerService.UpdateCareer(ctx, careerUpdateRequest, logo, cover)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success update career", nil)
}
