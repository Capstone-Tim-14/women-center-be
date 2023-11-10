package handlers

import (
	"net/http"
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type AuthHandler interface {
	UserAuthHandler(echo.Context) error
	OauthGoogleHandler(echo.Context) error
	OauthCallbackGoogleHandler(echo.Context) error
}

type AuthServiceImpl struct {
	AuthService services.AuthService
}

func NewAuthHandler(auth services.AuthService) AuthHandler {
	return &AuthServiceImpl{
		AuthService: auth,
	}
}

func (auth *AuthServiceImpl) UserAuthHandler(ctx echo.Context) error {

	var request requests.AuthRequest

	ErrBindRequest := ctx.Bind(&request)

	if ErrBindRequest != nil {
		return exceptions.BadRequestException("Invalid binding form input", ctx)
	}

	GetResponseAuth, ValidationErr, Err := auth.AuthService.UserAuthentication(request, ctx)

	if ValidationErr != nil {
		return exceptions.ValidationException(ctx, "Error validation", ValidationErr)
	}

	if Err != nil {
		if strings.Contains(Err.Error(), "Error uncorrect credential") {
			return exceptions.StatusUnauthorizedResponse(ctx, Err)
		}

		return exceptions.StatusInternalServerError(ctx, Err)
	}

	return responses.StatusOK(ctx, "Authentication Success", GetResponseAuth)
}

func (auth *AuthServiceImpl) OauthGoogleHandler(ctx echo.Context) error {

	GoogleAuth := auth.AuthService.GoogleAuthService()

	return ctx.Redirect(http.StatusSeeOther, GoogleAuth)
}
func (auth *AuthServiceImpl) OauthCallbackGoogleHandler(ctx echo.Context) error {

	GetResponseEmail, errMessage := auth.AuthService.GoogleCallbackService(ctx)

	if errMessage != nil {

		if strings.Contains(errMessage.Error(), "State is not match") {
			return exceptions.BadRequestException(errMessage.Error(), ctx)
		}

		if strings.Contains(errMessage.Error(), "User denied access login") {
			return exceptions.BadRequestException(errMessage.Error(), ctx)
		}

		return exceptions.BadRequestException(errMessage.Error(), ctx)
	}

	return responses.StatusCreated(ctx, "Google Authentication Success", GetResponseEmail)
}
