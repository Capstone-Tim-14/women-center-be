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
	DeleteScheduletById(ctx echo.Context, id int) error
	UpdateScheduleById(ctx echo.Context, request requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error)
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

		CheckSchedulingExists, _ := service.ScheduleRepo.FindStartEndDateCounseling(int(GetCounselorData.Id), val.Day_schedule, val.Time_start, val.Time_finish)

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

func (service *ScheduleServiceImpl) DeleteScheduletById(ctx echo.Context, id int) error {
	existingSchedule, _ := service.ScheduleRepo.FindById(id)
	if existingSchedule == nil {
		return fmt.Errorf("failed to find schedule")
	}

	err := service.ScheduleRepo.DeleteScheduleById(id)
	if err != nil {
		return fmt.Errorf("failed to remove schedule")
	}

	return nil
}

func (service *ScheduleServiceImpl) UpdateScheduleById(ctx echo.Context, request requests.CounselingScheduleRequest) ([]exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	scheduleID := ctx.Param("id")
	getId, errId := strconv.Atoi(scheduleID)
	if errId != nil {
		return nil, fmt.Errorf("invalid id")
	}

	Data, errData := service.ScheduleRepo.FindById(getId)
	if errData != nil {
		return nil, fmt.Errorf("schedule not found")
	}

	CounselorData, errCounselor := service.ScheduleRepo.GetSameDataCounselor(int(Data.Counselor_id))
	if errCounselor != nil {
		return nil, errCounselor
	}

	idSchedule := []int{}

	for _, data := range *CounselorData {
		idSchedule = append(idSchedule, int(data.Id))
	}

	sliceId := helpers.RemoveValue(idSchedule, getId)

	for _, id := range sliceId {

		OtherData, errOther := service.ScheduleRepo.FindById(id)
		if errOther != nil {
			return nil, errOther
		}

		if OtherData.Day_schedule == request.Day_schedule {
			if helpers.ParseClockToTime(request.Time_start).After(OtherData.Time_start) && helpers.ParseClockToTime(request.Time_start).Before(OtherData.Time_finish) {
				return nil, fmt.Errorf("schedule conflict with other")
			}
			if helpers.ParseClockToTime(request.Time_finish).Before(OtherData.Time_finish) && helpers.ParseClockToTime(request.Time_finish).After(OtherData.Time_start) {
				return nil, fmt.Errorf("schedule conflict with other")
			}
			if helpers.ParseClockToTime(request.Time_start).After(OtherData.Time_start) && helpers.ParseClockToTime(request.Time_finish).Before(OtherData.Time_finish) {
				return nil, fmt.Errorf("schedule conflict with other")
			}
			if helpers.ParseClockToTime(request.Time_start) == OtherData.Time_start {
				return nil, fmt.Errorf("schedule conflict with other")
			}
			if helpers.ParseClockToTime(request.Time_finish) == OtherData.Time_finish {
				return nil, fmt.Errorf("schedule conflict with other")
			}
		} else {
			continue
		}
	}

	schedule := reqConversion.ScheduleUpdateOnlyOne(request)

	err = service.ScheduleRepo.UpdateScheduleById(getId, &schedule)
	if err != nil {
		return nil, fmt.Errorf("error update schedule")
	}

	return nil, nil
}
