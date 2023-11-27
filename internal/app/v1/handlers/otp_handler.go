package handlers

import (
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type OtpHandler interface {
	SendOtpHandler(ctx echo.Context) error
}

type OtpHandlerImpl struct {
	OtpService services.OtpService
}

func NewOtpHandlerImpl(newOtpHandler OtpHandlerImpl) OtpHandler {
	return &newOtpHandler
}

func (handler *OtpHandlerImpl) SendOtpHandler(ctx echo.Context) error {

	var request requests.GenerateOTPRequest

	errBindingReq := ctx.Bind(&request)

	if errBindingReq != nil {
		return exceptions.BadRequestException(errBindingReq.Error(), ctx)
	}

	GetOTPCode, errGenerateCode := handler.OtpService.CreateAndSendingNewOtp(request)

	if errGenerateCode != nil {
		return exceptions.StatusInternalServerError(ctx, errGenerateCode)
	}

	return responses.StatusOK(ctx, "Success generate otp", GetOTPCode)
}
