package services

import (
	"fmt"
	"strconv"
	"woman-center-be/internal/app/v1/models/domain"
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
	GetUserLoginAndCounselorData(ctx echo.Context, counselor_id int) (*domain.Counselors, *domain.Users, error)
	CreateUserScheduleBooking(request requests.BookingCounselingRequest, counselor *domain.Counselors, userAuth *domain.Users) (*domain.UserScheduleCounseling, error)
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

func (service *BookingServiceImpl) CreateUserScheduleBooking(request requests.BookingCounselingRequest, counselor *domain.Counselors, userAuth *domain.Users) (*domain.UserScheduleCounseling, error) {
	GetDayRequestBooking := helpers.GetDayToTime(*helpers.ParseStringToTime(request.Booking_date))

	GetDayCounselorSchedule, errGetSchedule := service.CounselorRepo.FindCounselorAndGetOneOfSchedule(int(counselor.Id), GetDayRequestBooking)

	if errGetSchedule != nil {
		return nil, fmt.Errorf("One of schedule is not found")
	}

	CounselorRequestQuery := requests.UserScheduleCounselingQueryRequest{
		Counselor_schedule_id: GetDayCounselorSchedule.SingleSchedule.Id,
		Day:                   *helpers.ParseStringToTime(request.Booking_date),
		Time_start:            request.Booking_time,
	}

	fmt.Println(CounselorRequestQuery)

	IsExists, _ := service.ScheduleRepo.FindScheduleByDateAndTimeExist(CounselorRequestQuery)

	if IsExists {
		return nil, fmt.Errorf("One of schedule counselor is already booked to another user")
	}

	ConvertUserSchedule := conversion.UserScheduleCounselingToDomainSchedule(request, userAuth.Id, GetDayCounselorSchedule.SingleSchedule.Id)

	ResultUserSchedule, errCreateSchedule := service.ScheduleRepo.CreateUserScheduling(*userAuth, *ConvertUserSchedule)

	if errCreateSchedule != nil {
		return nil, errCreateSchedule
	}

	return ResultUserSchedule, nil
}

func (service *BookingServiceImpl) GetUserLoginAndCounselorData(ctx echo.Context, counselor_id int) (*domain.Counselors, *domain.Users, error) {
	GetCounselor, errGetData := service.CounselorRepo.FindById(counselor_id)

	if errGetData != nil {
		return nil, nil, errGetData
	}

	ClaimsUser := helpers.GetAuthClaims(ctx)
	GetUser, errGetUser := service.UserRepo.FindByID(int(ClaimsUser.Id))

	if errGetUser != nil {
		return nil, nil, fmt.Errorf("User profile not found")
	}

	return GetCounselor, GetUser, nil
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

	GetCounselor, GetUserAuth, errGetUser := service.GetUserLoginAndCounselorData(ctx, ConvertCounselorId)

	if errGetUser != nil {
		return nil, errGetUser
	}

	_, errGetPackage := service.CounselingPackage.FindById(ConvertCounselingId)

	if errGetPackage != nil {
		return nil, errGetPackage
	}

	var AllUserScheduleCounseling []domain.UserScheduleCounseling

	for _, val := range scheduleReq {

		DataUserSchedule, errCreateUserScheduleCounselingBook := service.CreateUserScheduleBooking(val, GetCounselor, GetUserAuth)

		if errCreateUserScheduleCounselingBook != nil {
			return nil, errCreateUserScheduleCounselingBook
		}

		AllUserScheduleCounseling = append(AllUserScheduleCounseling, *DataUserSchedule)

	}

	return nil, nil

}
