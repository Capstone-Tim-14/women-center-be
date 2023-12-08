package services

import (
	"fmt"
	"strconv"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type BookingService interface {
	CreateBookingCounseling(requests []requests.BookingCounselingRequest, ctx echo.Context) ([]exceptions.ValidationMessage, error)
}

type BookingServiceImpl struct {
	ScheduleRepo      repositories.UserScheduleCounselingRepository
	UserRepo          repositories.UserRepository
	CounselorRepo     repositories.CounselorRepository
	CounselingPackage repositories.CounselingPackageRepository
	Validate          *validator.Validate
}

func NewBookingService(bookingService BookingServiceImpl) BookingService {
	return &bookingService
}

func (service *BookingServiceImpl) CreateBookingCounseling(scheduleReq []requests.BookingCounselingRequest, ctx echo.Context) ([]exceptions.ValidationMessage, error) {

	for _, val := range scheduleReq {
		validationMessage := service.Validate.Struct(val)

		if validationMessage != nil {
			return helpers.ValidationError(ctx, validationMessage), nil
		}
	}

	ParamCounselorId := ctx.QueryParam("counselor_id")
	ParamCounselingPackageId := ctx.QueryParam("counseling_package_id")

	ConvertCounselorId, _ := strconv.Atoi(ParamCounselorId)
	ConvertCounselingId, _ := strconv.Atoi(ParamCounselingPackageId)

	GetCounselor, errGetData := service.CounselorRepo.FindById(ConvertCounselorId)

	if errGetData != nil {
		return nil, errGetData
	}

	ClaimsUser := helpers.GetAuthClaims(ctx)
	GetUser, errGetUser := service.UserRepo.FindByID(int(ClaimsUser.Id))

	if errGetUser != nil {
		return nil, fmt.Errorf("User profile not found")
	}

	_, errGetPackage := service.CounselingPackage.FindById(ConvertCounselingId)

	if errGetPackage != nil {
		return nil, errGetData
	}

	for _, val := range scheduleReq {

		GetDayRequestBooking := helpers.GetDayToTime(*helpers.ParseStringToTime(val.Booking_date))

		GetDayCounselorSchedule, errGetSchedule := service.CounselorRepo.FindCounselorAndGetOneOfSchedule(int(GetCounselor.Id), GetDayRequestBooking)

		if errGetSchedule != nil {
			return nil, fmt.Errorf("One of schedule is not found")
		}

		CounselorRequestQuery := requests.UserScheduleCounselingQueryRequest{
			Counselor_schedule_id: GetDayCounselorSchedule.SingleSchedule.Id,
			Day:                   *helpers.ParseStringToTime(val.Booking_date),
			Time_start:            val.Booking_time,
		}

		IsExists, _ := service.ScheduleRepo.FindScheduleByDateAndTimeExist(CounselorRequestQuery)

		if IsExists {
			return nil, fmt.Errorf("One of schedule counselor is already booked to another user")
		}

		ConvertUserSchedule := conversion.UserScheduleCounselingToDomainSchedule(val, GetUser.Id, GetDayCounselorSchedule.SingleSchedule.Id)

		_, errCreateSchedule := service.ScheduleRepo.CreateUserScheduling(*GetUser, *ConvertUserSchedule)

		if errCreateSchedule != nil {
			return nil, errCreateSchedule
		}

	}

	return nil, nil

}
