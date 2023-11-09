package services

import (
	"errors"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/pkg/oauth"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type AuthService interface {
	GoogleAuthService() string
	GoogleCallbackService(echo.Context) (*resources.UserGoogleInfo, error)
}

type AuthServiceImpl struct {
	validate *validator.Validate
}

func NewAuthService(validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		validate: validate,
	}
}

func (auth *AuthServiceImpl) GoogleAuthService() string {

	googleOauth := oauth.SetupGoogleOauth()
	url := googleOauth.AuthCodeURL(viper.GetString("GOOGLE_OAUTH.STATE_STRING"))

	return url

}

func (auth *AuthServiceImpl) GoogleCallbackService(ctx echo.Context) (*resources.UserGoogleInfo, error) {

	StateQuery := ctx.FormValue("state")
	CodeQuery := ctx.FormValue("code")

	if StateQuery != viper.GetString("GOOGLE_OAUTH.STATE_STRING") {
		return nil, errors.New("State is not match")
	}

	if CodeQuery == "" {
		return nil, errors.New("User denied access login")
	}

	googleSetup := oauth.SetupGoogleOauth()

	Response, ErrResponseGoogle := oauth.GetResponseAccountGoogle(CodeQuery, googleSetup)

	if ErrResponseGoogle != nil {
		return nil, errors.New("Error when processing google account")
	}

	return Response, nil

}
