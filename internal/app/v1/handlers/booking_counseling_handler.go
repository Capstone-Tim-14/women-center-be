package handlers

import (
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type BookingCounselingHandler interface {
	CreateBookingHandler(echo.Context) error
}

type BookingCounselingHandlerImpl struct {
	BookingService services.BookingService
}

func NewBookingCounselingHandler(bookingHandler BookingCounselingHandlerImpl) BookingCounselingHandler {
	return &bookingHandler
}

func (handler *BookingCounselingHandlerImpl) CreateBookingHandler(ctx echo.Context) error {

	var requests []requests.BookingCounselingRequest

	Binding := ctx.Bind(&requests)

	if Binding != nil {
		return exceptions.BadRequestException("Invalid format request", ctx)
	}

	validation, errBooking, result := handler.BookingService.CreateBookingCounseling(requests, ctx)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if errBooking != nil {
		return exceptions.StatusInternalServerError(ctx, errBooking)
	}

	return responses.StatusCreated(ctx, "Success created booking counseling", result)

}
