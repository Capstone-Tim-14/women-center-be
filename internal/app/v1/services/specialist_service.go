package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type SpecialistService interface {
	CreateSpecialist(ctx echo.Context, request requests.SpecialistRequest) (*domain.Specialist, []exceptions.ValidationMessage, error)
	GetListSpecialist(ctx echo.Context) ([]domain.Specialist, error)
	// DeleteSpecialistById(ctx echo.Context, id int) error
}

type SpecialistServiceImpl struct {
	SpecialistRepo repositories.SpecialistRepository
	validator      *validator.Validate
}

func NewSpecialistService(SpecialistRepo repositories.SpecialistRepository, validator *validator.Validate) SpecialistService {
	return &SpecialistServiceImpl{
		SpecialistRepo: SpecialistRepo,
		validator:      validator,
	}
}

func (service *SpecialistServiceImpl) CreateSpecialist(ctx echo.Context, request requests.SpecialistRequest) (*domain.Specialist, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	// Check if the tag already exists
	existingTag, _ := service.SpecialistRepo.FindSpecialistByName(request.Name)
	if existingTag != nil {
		return nil, nil, fmt.Errorf("Specialist name already exists")
	}

	newSpecialist := conversion.SpecialistCreateRequestToSpecialistDomain(request)

	// Create the new tag
	createdSpecialist, err := service.SpecialistRepo.CreateSpecialist(newSpecialist)
	if err != nil {
		return nil, nil, fmt.Errorf("Error while creating specialist: %s", err.Error())
	}

	return createdSpecialist, nil, nil
}

func (service *SpecialistServiceImpl) GetListSpecialist(ctx echo.Context) ([]domain.Specialist, error) {
	result, err := service.SpecialistRepo.FindAllSpecialist()

	if err != nil {
		return nil, err
	}

	return result, nil
}
