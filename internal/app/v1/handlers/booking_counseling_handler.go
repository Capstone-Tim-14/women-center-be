package handlers

import (
	"fmt"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	payment "woman-center-be/pkg/payment/midtrans"
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
	MidtransService    payment.MidtransCoreApi
}

func NewBookingCounselingHandler(bookingHandler BookingCounselingHandlerImpl) BookingCounselingHandler {
	return &bookingHandler
}

func (handler *BookingCounselingHandlerImpl) NotificationHandler(ctx echo.Context) error {

	notification := make(map[string]interface{})
	var errUpdateStatusBooking error

	errDecodeNotification := ctx.Bind(&notification)

	if errDecodeNotification != nil {
		return exceptions.StatusBadRequest(ctx, errDecodeNotification)
	}

	orderId, isExists := notification["order_id"].(string)

	if !isExists {
		return exceptions.StatusNotFound(ctx, fmt.Errorf("OrderId not found"))
	}

	GetTransactionStatus, errStatus := handler.MidtransService.CheckTransactionPayment(orderId)

	if errStatus != nil {
		return exceptions.StatusInternalServerError(ctx, errStatus)
	}

	if GetTransactionStatus == "challange" || GetTransactionStatus == "accept" || GetTransactionStatus == "settlement" || GetTransactionStatus == "deny" {
		_, errUpdateStatusBooking = handler.BookingService.UpdateStatusBooking(orderId, "SETTLEMENT")
	} else if GetTransactionStatus == "cancel" || GetTransactionStatus == "expire" {
		_, errUpdateStatusBooking = handler.BookingService.UpdateStatusBooking(orderId, "FAILED")
	} else if GetTransactionStatus == "pending" {
		return responses.StatusCreated(ctx, "Your transcation is pending, finish your payment", nil)
	}

	if errUpdateStatusBooking != nil {
		return exceptions.StatusInternalServerError(ctx, errUpdateStatusBooking)
	}

	return responses.StatusCreated(ctx, "Updated status transaction "+GetTransactionStatus, nil)

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
