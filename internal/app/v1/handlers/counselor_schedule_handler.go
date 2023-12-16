package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselorScheduleHandler interface {
	CreateScheduleHandler(echo.Context) error
	DeleteScheduleHandler(ctx echo.Context) error
	UpdateScheduleHandler(ctx echo.Context) error
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
		if strings.Contains(errCreate.Error(), "One of schedule is already exists") {
			return exceptions.StatusAlreadyExist(ctx, errCreate)
		}
		return exceptions.StatusInternalServerError(ctx, errCreate)
	}

	return responses.StatusCreated(ctx, "Schedule created", nil)
}

func (handler *ScheduleHandlerImpl) DeleteScheduleHandler(ctx echo.Context) error {
	idSchedule := ctx.Param("id")
	idInt, err := strconv.Atoi(idSchedule)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	err = handler.CounselorScheduleService.DeleteScheduletById(ctx, idInt)
	if err != nil {
		if strings.Contains(err.Error(), "failed to find schedule") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success remove schedule", nil)
}

func (handler *ScheduleHandlerImpl) UpdateScheduleHandler(ctx echo.Context) error {
	var request requests.CounselingScheduleRequest
	errBinding := ctx.Bind(&request)

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.CounselorScheduleService.UpdateScheduleById(ctx, request)
	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "invalid id") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "schedule not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Schedule updated!", nil)
}
