package services

import (
	"fmt"
	"mime/multipart"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/pkg/storage"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type EventService interface {
	CreateEvent(ctx echo.Context, request requests.EventRequest, poster *multipart.FileHeader) (*domain.Event, []exceptions.ValidationMessage, error)
	GetDetailEvent(ctx echo.Context, id int) (*domain.Event, error)
}

type EventServiceImpl struct {
	Validator *validator.Validate
	EventRepo repositories.EventRepository
}

func NewEventService(eventServiceImpl EventServiceImpl) EventService {
	return &eventServiceImpl
}

func (service *EventServiceImpl) CreateEvent(ctx echo.Context, request requests.EventRequest, poster *multipart.FileHeader) (*domain.Event, []exceptions.ValidationMessage, error) {
	PosterCloudURL, errUploadPoster := storage.S3PutFile(poster, "event/poster")

	if errUploadPoster != nil {
		return nil, nil, errUploadPoster
	}

	request.Poster = &PosterCloudURL
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	event := conversion.EventCreateRequestToEventDomain(&request)

	createEvent, err := service.EventRepo.CreateEvent(event)

	if err != nil {
		return nil, nil, fmt.Errorf("Error create event: %w", err)
	}

	return createEvent, nil, nil
}

func (service *EventServiceImpl) GetDetailEvent(ctx echo.Context, id int) (*domain.Event, error) {
	event, err := service.EventRepo.FindDetailEvent(id)

	if err != nil {
		return nil, fmt.Errorf("Error get detail event: %w", err)
	}

	return event, nil
}
