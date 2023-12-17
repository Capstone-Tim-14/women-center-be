package services

import (
	"fmt"
	"strconv"
	"time"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	resource "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type BookingService interface {
	CreateBookingCounseling(requests []requests.BookingCounselingRequest, ctx echo.Context) ([]exceptions.ValidationMessage, error, *resources.BookingCounselingResource)
	GetUserLoginAndCounselorData(ctx echo.Context, counselor_id int) (*domain.Counselors, *domain.Users, error)
	CreateUserScheduleBooking(request requests.BookingCounselingRequest, counselor *domain.Counselors, userAuth *domain.Users) (*domain.UserScheduleCounseling, error)
	UpdateStatusBooking(orderId string, status string) (bool, error)
	ListBookings(ctx echo.Context) ([]resources.BookingCounselingResource, error)
}

type BookingServiceImpl struct {
	UserService              UserService
	ScheduleRepo             repositories.UserScheduleCounselingRepository
	UserRepo                 repositories.UserRepository
	CounselorRepo            repositories.CounselorRepository
	CounselingPackage        repositories.CounselingPackageRepository
	CounselingPackageService CounselingPackageService
	BookingRepo              repositories.BookingCounselingRepository
	Validate                 *validator.Validate
}

func NewBookingService(bookingService BookingServiceImpl) BookingService {
	return &bookingService
}

func (service *BookingServiceImpl) ListBookings(ctx echo.Context) ([]resources.BookingCounselingResource, error) {

	GetUserClaim, ErrGetUser := service.UserService.GetUserProfile(ctx)

	if ErrGetUser != nil {
		return nil, ErrGetUser
	}

	BookingList, ErrBookingList := service.BookingRepo.GetBookingListByUser(GetUserClaim.Id)

	if ErrBookingList != nil {
		return nil, ErrBookingList
	}

	response := resource.ConvertBookingList(BookingList)

	return response, nil
}

func (service *BookingServiceImpl) UpdateStatusBooking(orderId string, status string) (bool, error) {

	ConvertToUUID, errConvert := uuid.FromString(orderId)

	if errConvert != nil {
		return false, fmt.Errorf("Error convert order id")
	}

	StatusUpdate, errUpdateBookingStatus := service.BookingRepo.UpdateStatusBooking(ConvertToUUID, status)

	if errUpdateBookingStatus != nil {
		return false, fmt.Errorf("failed update")
	}

	return StatusUpdate, nil
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

	GetUser, errGetUser := service.UserService.GetUserProfile(ctx)

	if errGetUser != nil {
		return nil, nil, fmt.Errorf("User profile not found")
	}

	return GetCounselor, GetUser, nil
}

func (service *BookingServiceImpl) CreateBookingCounseling(scheduleReq []requests.BookingCounselingRequest, ctx echo.Context) ([]exceptions.ValidationMessage, error, *resources.BookingCounselingResource) {

	for _, val := range scheduleReq {
		validationMessage := service.Validate.Struct(val)

		if validationMessage != nil {
			return helpers.ValidationError(ctx, validationMessage), nil, nil
		}
	}

	ParamCounselorId := ctx.QueryParam("counselor_id")
	ParamCounselingPackageId := ctx.QueryParam("counseling_package_id")

	ConvertCounselorId, _ := strconv.Atoi(ParamCounselorId)
	ConvertCounselingId, _ := strconv.Atoi(ParamCounselingPackageId)

	GetCounselor, GetUserAuth, errGetUser := service.GetUserLoginAndCounselorData(ctx, ConvertCounselorId)

	if errGetUser != nil {
		return nil, errGetUser, nil
	}

	var AllUserScheduleCounseling []domain.UserScheduleCounseling

	for _, val := range scheduleReq {

		DataUserSchedule, errCreateUserScheduleCounselingBook := service.CreateUserScheduleBooking(val, GetCounselor, GetUserAuth)

		if errCreateUserScheduleCounselingBook != nil {
			return nil, errCreateUserScheduleCounselingBook, nil
		}

		AllUserScheduleCounseling = append(AllUserScheduleCounseling, *DataUserSchedule)

	}

	ConvertCounselingPackage, errGetCounselingPackage := service.CounselingPackageService.GetPackageById(ctx, ConvertCounselingId)

	if errGetCounselingPackage != nil {
		return nil, errGetCounselingPackage, nil
	}

	transactionReq := requests.TransactionBookingRequest{}

	transactionReq.User_id = GetUserAuth.Id
	transactionReq.Transaction_date = time.Now()
	transactionReq.Status = "IN PROCESS"
	transactionReq.Transaction_detail.Counseling_package_id = ConvertCounselingPackage.Id
	transactionReq.Transaction_detail.Tax = helpers.StringToDecimal("8000")
	transactionReq.Transaction_detail.Total = helpers.TotalTransaction(transactionReq.Transaction_detail.Tax, ConvertCounselingPackage.Price)

	ConvertBookingCounseling := conversion.BookingCounselingRequestToDomain(transactionReq)

	resultTransaction, errCreateTransaction := service.BookingRepo.CreateBooking(&ConvertBookingCounseling)

	if errCreateTransaction != nil {
		return nil, fmt.Errorf("Failed to create transaction"), nil
	}

	UpdateBookingDetailToSchedule := service.ScheduleRepo.UpdateMultipleScheduleBooked(AllUserScheduleCounseling, resultTransaction.Booking_detail_id)

	if UpdateBookingDetailToSchedule != nil {
		return nil, fmt.Errorf("Failed processing transaction"), nil
	}

	result := resource.BookingCounselingToDomainBookingCounseling(resultTransaction, ConvertCounselingPackage, GetUserAuth, GetCounselor, AllUserScheduleCounseling)

	return nil, nil, result

}
