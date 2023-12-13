package handlers

import (
	"fmt"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
	uuid "github.com/satori/go.uuid"
)

type BookingCounselingHandler interface {
	CreateBookingHandler(echo.Context) error
	CreateTransactionPaymentHandler(echo.Context) error
	NotificationHandler(echo.Context) error
}

type BookingCounselingHandlerImpl struct {
	BookingService     services.BookingService
	TransactionService services.TransactionPaymentService
}

func NewBookingCounselingHandler(bookingHandler BookingCounselingHandlerImpl) BookingCounselingHandler {
	return &bookingHandler
}

func (handler *BookingCounselingHandlerImpl) NotificationHandler(ctx echo.Context) error {

	notification := make(map[string]interface{})

	errDecodeNotification := ctx.Bind(&notification)

	if errDecodeNotification != nil {
		return exceptions.StatusBadRequest(ctx, errDecodeNotification)
	}

	orderId, isExists := notification["order_id"].(string)

	if !isExists {
		return exceptions.StatusNotFound(ctx, fmt.Errorf("OrderId not found"))
	}

}

func (handler *BookingCounselingHandlerImpl) CreateTransactionPaymentHandler(ctx echo.Context) error {

	orderId := ctx.QueryParam("order_id")
	paymentCode := ctx.QueryParam("payment_code")

	if orderId == "" {
		return exceptions.BadRequestException("order id required", ctx)
	}

	if paymentCode == "" {
		return exceptions.BadRequestException("payment code", ctx)
	}

	parseOrderId, errParse := uuid.FromString(orderId)

	if errParse != nil {
		return exceptions.BadRequestException("invalid format order id", ctx)
	}

	ResultTransaction, errTransaction := handler.TransactionService.TransactionPaymentService(parseOrderId, paymentCode)

	if errTransaction != nil {
		return exceptions.StatusInternalServerError(ctx, errTransaction)
	}

	return responses.StatusCreated(ctx, "Invoice created successfully, please complete your payment", ResultTransaction)

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
