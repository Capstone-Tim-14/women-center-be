package services

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ScheduleService interface {
	CreateSchedule(ctx echo.Context, request requests.CounselingScheduleRequest) (*domain.Counseling_Schedule, []exceptions.ValidationMessage, error)
}

type ScheduleServiceImpl struct {
	ScheduleRepo  repositories.ScheduleRepository
	Validator     *validator.Validate
	CounselorRepo repositories.CounselorRepository
}

func NewScheduleService(scheduleRepo repositories.ScheduleRepository, validator *validator.Validate, counselorRepo repositories.CounselorRepository) ScheduleService {
	return &ScheduleServiceImpl{
		ScheduleRepo:  scheduleRepo,
		Validator:     validator,
		CounselorRepo: counselorRepo,
	}
}

func (service *ScheduleServiceImpl) CreateSchedule(ctx echo.Context, request requests.CounselingScheduleRequest) (*domain.Counseling_Schedule, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingSchedule, _ := service.ScheduleRepo.FindFreeSchedule(request.Day_schedule, request.Time_start, request.Time_finish)

	// NOT YET
}
