package handlers

import (
	"net/http"
	"woman-center-be/internal/app/v1/services"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type AuthHandler interface {
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

func (auth *AuthServiceImpl) OauthGoogleHandler(ctx echo.Context) error {

	GoogleAuth := auth.AuthService.GoogleAuthService()

	return ctx.Redirect(http.StatusSeeOther, GoogleAuth)
}
func (auth *AuthServiceImpl) OauthCallbackGoogleHandler(ctx echo.Context) error {

	Getquery := ctx.Request().URL.Query()

	if Getquery["state"][0] != viper.GetString("GOOGLE_OAUTH.STATE_STRING") {
		return echo.NewHTTPError(http.StatusInternalServerError, "state dont match")
	}

	return nil
}
