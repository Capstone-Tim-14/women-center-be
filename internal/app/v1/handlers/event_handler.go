package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type EventHandler interface {
	CreateEvent(ctx echo.Context) error
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
