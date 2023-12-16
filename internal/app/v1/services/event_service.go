package services

import (
	"fmt"
	"mime/multipart"
	"strconv"
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
	GetAllEvent(ctx echo.Context) ([]domain.Event, error)
	GetAllEventMobile(ctx echo.Context) ([]domain.Event, error)
	UpdateEvent(ctx echo.Context, request requests.EventRequest, poster *multipart.FileHeader) ([]exceptions.ValidationMessage, error)
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

	event := conversion.EventCreateRequestToEventDomain(request)

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

func (service *EventServiceImpl) GetAllEvent(ctx echo.Context) ([]domain.Event, error) {
	events, err := service.EventRepo.FindAllEvent()

	if err != nil {
		return nil, fmt.Errorf("Error get all event: %w", err)
	}

	return events, nil
}

func (service *EventServiceImpl) GetAllEventMobile(ctx echo.Context) ([]domain.Event, error) {
	events, err := service.EventRepo.FindAllEvent()

	if err != nil {
		return nil, fmt.Errorf("Error get all event: %w", err)
	}

	return events, nil
}

func (service *EventServiceImpl) UpdateEvent(ctx echo.Context, request requests.EventRequest, poster *multipart.FileHeader) ([]exceptions.ValidationMessage, error) {

	if poster != nil {

		PosterCloudURL, errUploadPoster := storage.S3PutFile(poster, "event/poster")

		if errUploadPoster != nil {
			return nil, errUploadPoster
		}

		request.Poster = &PosterCloudURL
	}

	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	id := ctx.Param("id")
	getId, _ := strconv.Atoi(id)

	_, err = service.EventRepo.FindById(getId)
	if err != nil {
		return nil, fmt.Errorf("Event not found")

	}

	event := conversion.EventCreateRequestToEventDomain(request)

	_, err = service.EventRepo.UpdateEvent(getId, event), nil
	if err != nil {
		return nil, fmt.Errorf("Error update event")
	}

	return nil, nil
}
