package services

import (
	"errors"
	"strings"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	conversionResource "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/pkg/oauth"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type AuthService interface {
	UserAuthentication(requests.AuthRequest, echo.Context) (*resources.AuthTokenResource, []exceptions.ValidationMessage, error)
	GoogleAuthService() string
	GoogleCallbackService(echo.Context) (*resources.AuthTokenResource, error)
}

type AuthServiceImpl struct {
	RoleRepo repositories.RoleRepository
	UserRepo repositories.UserRepository
	validate *validator.Validate
}

func NewAuthService(role repositories.RoleRepository, user repositories.UserRepository, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		RoleRepo: role,
		UserRepo: user,
		validate: validate,
	}
}

func (auth *AuthServiceImpl) GoogleAuthService() string {

	googleOauth := oauth.SetupGoogleOauth()
	url := googleOauth.AuthCodeURL(viper.GetString("GOOGLE_OAUTH.STATE_STRING"))

	return url

}

func (auth *AuthServiceImpl) GoogleCallbackService(ctx echo.Context) (*resources.AuthTokenResource, error) {
	var SetAuthenticateData resources.AuthResource
	var GetAuthWithTokenResponse *resources.AuthTokenResource

	StateQuery := ctx.FormValue("state")
	CodeQuery := ctx.FormValue("code")

	if StateQuery != viper.GetString("GOOGLE_OAUTH.STATE_STRING") {
		return nil, errors.New("State is not match")
	}

	if CodeQuery == "" {
		return nil, errors.New("User denied access login")
	}

	googleSetup := oauth.SetupGoogleOauth()

	UserGoogleResponse, ErrResponseGoogle := oauth.GetResponseAccountGoogle(CodeQuery, googleSetup)

	if ErrResponseGoogle != nil {
		return nil, errors.New("Error when processing google account")
	}

	UserResponse, ErrExists := auth.UserRepo.FindyByEmail(UserGoogleResponse.Email)

	if ErrExists != nil {

		SetUser := requests.UserRequest{
			First_name:      UserGoogleResponse.GivenName,
			Last_name:       UserGoogleResponse.FamilyName,
			Email:           UserGoogleResponse.Email,
			Profile_picture: UserGoogleResponse.Picture,
			Address:         UserGoogleResponse.Locale,
			Username:        strings.ToLower(UserGoogleResponse.Name),
			Password:        helpers.StringWithCharset(10),
			Phone_number:    62123456789,
		}

		GetRole, _ := auth.RoleRepo.FindByName("user")

		if GetRole == nil {
			return nil, errors.New("Role not found")
		}

		SetUser.Role_id = uint(GetRole.Id)

		UserConvert := conversion.UserCreateRequestToUserDomain(SetUser)
		UserConvert.Credential.Password = helpers.HashPassword(SetUser.Password)

		UserCreate, ErrCreate := auth.UserRepo.CreateUser(UserConvert)

		if ErrCreate != nil {
			return nil, errors.New("Failed when processing user data")
		}

		GetUserByEmail, ErrGetUser := auth.UserRepo.FindyByEmail(UserCreate.Email)

		if ErrGetUser != nil {
			return nil, errors.New("Failed when processing user data")
		}

		SetAuthenticateData = conversionResource.UserDomainToAuthResource(GetUserByEmail)

	} else {
		SetAuthenticateData = conversionResource.UserDomainToAuthResource(UserResponse)
	}

	GetTokenAuth, ErrGetToken := helpers.GenerateToken(SetAuthenticateData, ctx)

	if ErrGetToken != nil {
		return nil, ErrGetToken
	}

	GetAuthWithTokenResponse = conversionResource.AuthResourceToAuthTokenResource(SetAuthenticateData, GetTokenAuth)

	return GetAuthWithTokenResponse, nil

}

func (service *AuthServiceImpl) UserAuthentication(request requests.AuthRequest, ctx echo.Context) (*resources.AuthTokenResource, []exceptions.ValidationMessage, error) {

	ValidationErr := service.validate.Struct(request)

	if ValidationErr != nil {
		return nil, helpers.ValidationError(ctx, ValidationErr), nil
	}

	CheckUserAuthentication, UserErr := service.UserRepo.FindyByEmail(request.Email)

	if UserErr != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	ErrComparePassword := helpers.ComparePassword(CheckUserAuthentication.Credential.Password, request.Password)

	if ErrComparePassword != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	UserConvert := conversionResource.UserDomainToAuthResource(CheckUserAuthentication)

	GetToken, ErrToken := helpers.GenerateToken(UserConvert, ctx)

	if ErrToken != nil {
		return nil, nil, errors.New("Failed generate token")
	}

	GetAuthResponse := conversionResource.AuthResourceToAuthTokenResource(UserConvert, GetToken)

	return GetAuthResponse, nil, nil

}
