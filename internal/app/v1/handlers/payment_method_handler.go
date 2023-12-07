package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type PaymentMethodHandler interface {
	ListPaymentMethodHandler(ctx echo.Context) error
}

type PaymentMethodHandlerImpl struct {
	PaymentMethodService services.PaymentMethodService
}

func NewPaymentMethodHandler(paymentMethodImpl PaymentMethodHandlerImpl) PaymentMethodHandler {
	return &paymentMethodImpl
}

func (handler *PaymentMethodHandlerImpl) ListPaymentMethodHandler(ctx echo.Context) error {

	GetlistBankMethod, errGetList := handler.PaymentMethodService.GetListBankMethod()

	if errGetList != nil {
		if strings.Contains(errGetList.Error(), "Bank method is empty") {
			return exceptions.StatusNotFound(ctx, errGetList)
		}
		return exceptions.StatusInternalServerError(ctx, errGetList)
	}

	return responses.StatusOK(ctx, "Success get payment bank", GetlistBankMethod)
}
