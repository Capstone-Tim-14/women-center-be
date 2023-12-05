package services

import (
	"fmt"
	"strconv"
	"woman-center-be/internal/app/v1/repositories"
	reqConversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ScheduleService interface {
	CreateSchedule(ctx echo.Context, request []requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error)
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

func (service *ScheduleServiceImpl) CreateSchedule(ctx echo.Context, requests []requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error) {

	for _, val := range requests {

		err := service.Validator.Struct(val)

		if err != nil {
			return helpers.ValidationError(ctx, err), nil
		}

	}

	GetId := ctx.Param("id")
	ParseToInt, errParse := strconv.Atoi(GetId)

	if errParse != nil {
		return nil, fmt.Errorf("Invalid format integer")
	}

	GetCounselorData, errGetCounselor := service.CounselorRepo.FindById(ParseToInt)

	if errGetCounselor != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	SchedulingConvert := reqConversion.CounselorScheduleToCounselorDomain(requests)

	for _, val := range SchedulingConvert {

		CheckSchedulingExists, _ := service.ScheduleRepo.CheckDayCounselingScheduleExists(int(GetCounselorData.Id), val.Day_schedule)

		if CheckSchedulingExists != nil {
			return nil, fmt.Errorf("One of schedule is already exists")
		}

	}

	errCreate := service.ScheduleRepo.CreateSchedule(GetCounselorData, SchedulingConvert)

	if errCreate != nil {
		return nil, fmt.Errorf("Error when create schedule")
	}

	return nil, nil
}
