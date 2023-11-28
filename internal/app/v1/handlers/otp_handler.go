package handlers

import (
	"fmt"
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type OtpHandler interface {
	SendOtpHandler(ctx echo.Context) error
	VerifyTokenHandler(ctx echo.Context) error
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

	errGenerateCode := handler.OtpService.CreateAndSendingNewOtp(request)

	if errGenerateCode != nil {
		return exceptions.StatusInternalServerError(ctx, errGenerateCode)
	}

	return responses.StatusOK(ctx, "Success send otp to your email.", nil)
}

func (handler *OtpHandlerImpl) VerifyTokenHandler(ctx echo.Context) error {
	var request requests.VerifyOTPRequest

	ErrBinding := ctx.Bind(&request)

	if ErrBinding != nil {
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Error no request send"))
	}

	ErrVerifyOTP := handler.OtpService.VerifyOTP(request)

	if ErrVerifyOTP != nil {

		if strings.Contains(ErrVerifyOTP.Error(), "Invalid OTP or user dosen't exists") {
			return exceptions.StatusBadRequest(ctx, ErrVerifyOTP)
		}

		return exceptions.StatusInternalServerError(ctx, ErrVerifyOTP)
	}

	return responses.StatusOK(ctx, "OTP Verification Complete", nil)

}
