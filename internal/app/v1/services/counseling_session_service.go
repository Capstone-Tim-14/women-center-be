package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/resources/v1"

	"github.com/labstack/echo/v4"
)

type CounselingSessionService interface {
	GetListCounselingSession(echo.Context) ([]resources.CounselingSessioningResource, error)
}

type CounselingSessionServiceImpl struct {
	CounselorRepo    repositories.CounselorRepository
	BookingRepo      repositories.BookingCounselingRepository
	CounselorService CounselorService
}

func NewCounselingSessionService(counselingSession CounselingSessionServiceImpl) CounselingSessionService {
	return &counselingSession
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
