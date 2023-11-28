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

type CounselorHandler interface {
	RegisterHandler(ctx echo.Context) error
	AddSpecialist(ctx echo.Context) error
	DeleteSpecialist(ctx echo.Context) error
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

func (handler *CounselorHandlerImpl) AddSpecialist(ctx echo.Context) error {
	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)
	var request requests.CounselorHasSpecialistRequest
	errBinding := ctx.Bind(&request)
	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.CounselorService.AddSpecialist(ctx, uint(convertid), request)
	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success add specialist to counselor", nil)
}

func (handler *CounselorHandlerImpl) DeleteSpecialist(ctx echo.Context) error {
	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)
	var request requests.DeleteCounselorSpecialist
	errBinding := ctx.Bind(&request)
	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.CounselorService.DeleteSpecialist(ctx, uint(convertid), request)
	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success delete specialist counselor", nil)
}
