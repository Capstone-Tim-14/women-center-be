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

type CounselorHandler interface {
	RegisterHandler(echo.Context) error
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
		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorCreateResponse := conversion.CounselorDomainToCounselorResponse(response)

	return responses.StatusCreated(ctx, "Counselor created successfully", counselorCreateResponse)

}

// Not Yet
