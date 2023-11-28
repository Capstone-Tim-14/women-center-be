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

type JobTypeService interface {
	CreateJobType(ctx echo.Context, request requests.JobTypeRequest) (*domain.Job_Type, []exceptions.ValidationMessage, error)
	GetJobType(ctx echo.Context) ([]domain.Job_Type, error)
}

type JobTypeServiceImpl struct {
	JobTypeRepo repositories.JobTypeRepository
	validator   *validator.Validate
}

func NewJobTypeService(jobtypeRepo repositories.JobTypeRepository, validator *validator.Validate) JobTypeService {
	return &JobTypeServiceImpl{
		JobTypeRepo: jobtypeRepo,
		validator:   validator,
	}
}

func (service *JobTypeServiceImpl) CreateJobType(ctx echo.Context, request requests.JobTypeRequest) (*domain.Job_Type, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingTypes, _ := service.JobTypeRepo.FindJobTypeByName(request.Name)
	if existingTypes != nil {
		return nil, nil, fmt.Errorf("job type already exists")
	}

	newJobType := conversion.JobTypeCreateRequestToJobTypeDomain(request)

	// Create the new tag
	createJobType, err := service.JobTypeRepo.CreateJobType(newJobType)
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating job type: %s", err.Error())
	}

	return createJobType, nil, nil
}

func (service *JobTypeServiceImpl) GetJobType(ctx echo.Context) ([]domain.Job_Type, error) {
	result, err := service.JobTypeRepo.ShowAllJobType()

	if err != nil {
		return nil, err
	}

	return result, nil
}
