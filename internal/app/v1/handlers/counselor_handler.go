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

type CounselorHandler interface {
	RegisterHandler(echo.Context) error
	GetAllCounselorsHandler(echo.Context) error
	UpdateCounselorHandler(echo.Context) error
}

type CounselorHandlerImpl struct {
	CounselorService services.CounselorService
}

func NewCounselorHandler(counselor services.CounselorService) CounselorHandler {
	return &CounselorHandlerImpl{
		CounselorService: counselor,
	}
}

func (handler *CounselorHandlerImpl) RegisterHandler(ctx echo.Context) error {
	counselorCreateRequest := requests.CounselorRequest{}
	err := ctx.Bind(&counselorCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.CounselorService.RegisterCounselor(ctx, counselorCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exist") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorCreateResponse := conversion.CounselorDomainToCounselorResponse(response)

	return responses.StatusCreated(ctx, "Counselor created successfully", counselorCreateResponse)

}

func (handler *CounselorHandlerImpl) GetAllCounselorsHandler(ctx echo.Context) error {
	response, err := handler.CounselorService.GetAllCounselors(ctx)

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorResponse := conversion.ConvertCounselorDomainToCounselorResponse(response)

	return responses.StatusOK(ctx, "Get all counselors successfully", counselorResponse)
}

func (handler *CounselorHandlerImpl) UpdateCounselorHandler(ctx echo.Context) error {
	counselorUpdateRequest := requests.CounselorRequest{}
	picture, _ := ctx.FormFile("picture")
	err := ctx.Bind(&counselorUpdateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.CounselorService.UpdateCounselor(ctx, counselorUpdateRequest, picture)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Counselor updated successfully", nil)
}
