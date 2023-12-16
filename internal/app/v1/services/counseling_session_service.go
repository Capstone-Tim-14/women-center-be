package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/resources/v1"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type CounselingSessionService interface {
	GetListCounselingSession(echo.Context) ([]resources.CounselingSessioningResource, error)
	GetCounselingSessionDetail(echo.Context) (*resources.CounselingSessionDetailResource, error)
}

type CounselingSessionServiceImpl struct {
	CounselorRepo    repositories.CounselorRepository
	BookingRepo      repositories.BookingCounselingRepository
	CounselorService CounselorService
}

func NewCounselingSessionService(counselingSession CounselingSessionServiceImpl) CounselingSessionService {
	return &counselingSession
}

func (service *CounselingSessionServiceImpl) GetCounselingSessionDetail(ctx echo.Context) (*resources.CounselingSessionDetailResource, error) {

	paramOrderId := ctx.Param("order_id")

	convertOrderId, errConvert := uuid.FromString(paramOrderId)

	if errConvert != nil {
		return nil, fmt.Errorf("invalid format order_id")
	}

	CounselorProfile, errCounselorProfile := service.CounselorService.GetCounselorProfile(ctx)

	if errCounselorProfile != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	CounselorSessionDetail, CounselingScheduleSessions, errGetCounseling := service.BookingRepo.GetBookingCounselingDetail(CounselorProfile.Id, convertOrderId)

	if errGetCounseling != nil {
		return nil, errGetCounseling
	}

	Response := conversion.CounselingSessionBookedDetailConvert(CounselorSessionDetail, CounselingScheduleSessions)

	return &Response, nil

}

func (service *CounselingSessionServiceImpl) GetListCounselingSession(ctx echo.Context) ([]resources.CounselingSessioningResource, error) {

	CounselorProfile, errCounselorProfile := service.CounselorService.GetCounselorProfile(ctx)

	if errCounselorProfile != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	GetUserCounselingBooked, errGetBooking := service.BookingRepo.GetBookingListByCounselor(CounselorProfile.Id)

	if errGetBooking != nil {
		return nil, errGetBooking
	}

	setResponseCounselingBooked := conversion.CounselingSessionBookedConvert(GetUserCounselingBooked)

	return setResponseCounselingBooked, nil

}
