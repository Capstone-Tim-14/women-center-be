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

type JobTypeHandler interface {
	CreateJobType(ctx echo.Context) error
	ShowAllJobType(ctx echo.Context) error
}

type JobTypeHandlerImpl struct {
	JobTypeService services.JobTypeService
}

func NewJobTypeHandler(jobtypeService services.JobTypeService) JobTypeHandler {
	return &JobTypeHandlerImpl{
		JobTypeService: jobtypeService,
	}
}

func (handler *JobTypeHandlerImpl) CreateJobType(ctx echo.Context) error {
	jobtypeCreateRequest := requests.JobTypeRequest{}
	err := ctx.Bind(&jobtypeCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.JobTypeService.CreateJobType(ctx, jobtypeCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {

		if strings.Contains(err.Error(), "job type already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	jobtypeCreateResponse := conversion.JobTypeDomainToJobTypeResponse(response)

	return responses.StatusCreated(ctx, "Job type created successfully", jobtypeCreateResponse)
}

func (handler *JobTypeHandlerImpl) ShowAllJobType(ctx echo.Context) error {
	response, err := handler.JobTypeService.GetJobType(ctx)
	if err != nil {

		if strings.Contains(err.Error(), "Job types is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	jobtypeResponse := conversion.ConvertJobTypeResource(response)

	return responses.StatusOK(ctx, "Success Get All Job Types", jobtypeResponse)

}
