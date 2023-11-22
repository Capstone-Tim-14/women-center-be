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

type SpecialistHandler interface {
	CreateSpecialist(ctx echo.Context) error
	GetListSpecialistHandler(ctx echo.Context) error
}

type SpecialistHandlerImpl struct {
	SpecialistService services.SpecialistService
}

func NewSpecialistHandler(specialistService services.SpecialistService) SpecialistHandler {
	return &SpecialistHandlerImpl{
		SpecialistService: specialistService,
	}
}

func (handler *SpecialistHandlerImpl) CreateSpecialist(ctx echo.Context) error {
	specialistCreateRequest := requests.SpecialistRequest{}
	err := ctx.Bind(&specialistCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	// Call TagService to create a new tag
	response, validation, err := handler.SpecialistService.CreateSpecialist(ctx, specialistCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {

		if strings.Contains(err.Error(), "specialist name already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	// Prepare the response
	specialistCreateResponse := conversion.SpecialistDomainToSpecialistResponse(response)

	return responses.StatusCreated(ctx, "Specialist name created successfully", specialistCreateResponse)
}

func (handler *SpecialistHandlerImpl) GetListSpecialistHandler(ctx echo.Context) error {
	response, err := handler.SpecialistService.GetListSpecialist(ctx)
	if err != nil {

		if strings.Contains(err.Error(), "specialist name is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	specialistResponse := conversion.ConvertSpecialistResource(response)

	return responses.StatusOK(ctx, "Success Get All Specialist", specialistResponse)

}
