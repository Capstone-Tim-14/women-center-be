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

type EventHandler interface {
	CreateEvent(ctx echo.Context) error
	GetDetailEvent(ctx echo.Context) error
	GetDetailEventMobile(ctx echo.Context) error
	GetAllEvent(ctx echo.Context) error
	GetAllEventMobile(ctx echo.Context) error
}

type EventHandlerImpl struct {
	EventService services.EventService
}

func NewEventHandler(event services.EventService) EventHandler {
	return &EventHandlerImpl{
		EventService: event,
	}

}

func (handler *EventHandlerImpl) CreateEvent(ctx echo.Context) error {
	Eventrequest := requests.EventRequest{}
	Poster, errPoster := ctx.FormFile("poster")

	if errPoster != nil {
		return exceptions.StatusBadRequest(ctx, errPoster)
	}

	err := ctx.Bind(&Eventrequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.EventService.CreateEvent(ctx, Eventrequest, Poster)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Error create event") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success Create Event", nil)
}

func (handler *EventHandlerImpl) GetDetailEvent(ctx echo.Context) error {
	getId := ctx.Param("id")
	eventId, _ := strconv.Atoi(getId)

	event, err := handler.EventService.GetDetailEvent(ctx, eventId)

	if err != nil {
		if strings.Contains(err.Error(), "Error get detail event") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	result := conversion.EventDetailDomainToEventResource(event)

	return responses.StatusOK(ctx, "Success Get Detail Event", result)
}

func (handler *EventHandlerImpl) GetDetailEventMobile(ctx echo.Context) error {
	getId := ctx.Param("id")
	eventId, _ := strconv.Atoi(getId)

	event, err := handler.EventService.GetDetailEvent(ctx, eventId)

	if err != nil {
		if strings.Contains(err.Error(), "Error get detail event") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	result := conversion.EventDetailDomainToEventResource(event)

	return responses.StatusOK(ctx, "Success Get Detail Event", result)
}

func (handler *EventHandlerImpl) GetAllEvent(ctx echo.Context) error {
	events, err := handler.EventService.GetAllEvent(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Error get all event") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	result := conversion.AllEventConvertResource(events)

	return responses.StatusOK(ctx, "Success Get All Event", result)
}

func (handler *EventHandlerImpl) GetAllEventMobile(ctx echo.Context) error {
	events, err := handler.EventService.GetAllEvent(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Error get all event") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	result := conversion.AllEventConvertResource(events)

	return responses.StatusOK(ctx, "Success Get All Event", result)
}
