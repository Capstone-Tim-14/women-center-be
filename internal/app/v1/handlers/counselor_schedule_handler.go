package handlers

import (
	"fmt"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselorScheduleHandler interface {
	CreateScheduleHandler(echo.Context) error
}

type ScheduleHandlerImpl struct {
	CounselorService         services.CounselorService
	CounselorScheduleService services.ScheduleService
}

func NewCounselorScheduleHandler(CounselorImpl ScheduleHandlerImpl) CounselorScheduleHandler {
	return &CounselorImpl
}

func (handler *ScheduleHandlerImpl) CreateScheduleHandler(ctx echo.Context) error {

	var requests []requests.CounselingScheduleRequest

	Binding := ctx.Bind(&requests)

	if Binding != nil {
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Error invalid format requests"))
	}

	validation, errCreate := handler.CounselorScheduleService.CreateSchedule(ctx, requests)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Validation requests", validation)
	}

	if errCreate != nil {
		return exceptions.StatusInternalServerError(ctx, errCreate)
	}

	return responses.StatusCreated(ctx, "Schedule created", nil)
}
