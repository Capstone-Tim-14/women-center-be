package services

import (
	"woman-center-be/pkg/oauth"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

type AuthService interface {
	GoogleAuthService() string
}

type AuthServiceImpl struct {
	GoogleOauthConfig *oauth2.Config
	validate          *validator.Validate
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
