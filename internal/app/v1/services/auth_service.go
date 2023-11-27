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
	AdminAuthentication(requests.AuthRequest, echo.Context) (*resources.AuthTokenResource, []exceptions.ValidationMessage, error)
	GoogleAuthService() string
	GoogleCallbackService(echo.Context) (*resources.AuthTokenResource, error)
}

type AuthServiceImpl struct {
	AdminRepo      repositories.AdminRepository
	RoleRepo       repositories.RoleRepository
	UserRepo       repositories.UserRepository
	CounselorRepo  repositories.CounselorRepository
	CredentialRepo repositories.CredentialRepository
	Validate       *validator.Validate
}

func NewAuthService(authServiceImpl AuthServiceImpl) AuthService {
	return &authServiceImpl
}

func (auth *AuthServiceImpl) AdminAuthentication(request requests.AuthRequest, ctx echo.Context) (*resources.AuthTokenResource, []exceptions.ValidationMessage, error) {

	ValidationErr := auth.Validate.Struct(request)

	if ValidationErr != nil {
		return nil, helpers.ValidationError(ctx, ValidationErr), nil
	}

	CheckAdminEmail, ErrCheckEmail := auth.AdminRepo.FindyByEmail(request.Email)

	if ErrCheckEmail != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	ErrComparePassword := helpers.ComparePassword(CheckAdminEmail.Credential.Password, request.Password)

	if ErrComparePassword != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	AdminConvert := conversionResource.AdminDomainToAuthResource(CheckAdminEmail)

	GetToken, ErrToken := helpers.GenerateToken(AdminConvert, ctx)

	if ErrToken != nil {
		return nil, nil, errors.New("Failed generate token")
	}

	GetAuthResponse := conversionResource.AuthResourceToAuthTokenResource(AdminConvert, GetToken)

	return GetAuthResponse, nil, nil

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
			Username:        strings.ToLower(UserGoogleResponse.Name),
			Password:        helpers.StringWithCharset(10),
			Phone_number:    "62123456789",
		}

		GetRole, _ := auth.RoleRepo.FindByName("user")

		if GetRole == nil {
			return nil, errors.New("Role not found")
		}

		SetUser.Role_id = uint(GetRole.Id)

		UserConvert := conversion.UserCreateRequestToUserDomain(SetUser)
		UserConvert.Credential.Password = helpers.HashPassword(SetUser.Password)

		_, ErrCreate := auth.UserRepo.CreateUser(UserConvert)

		if ErrCreate != nil {
			return nil, errors.New("Failed when processing user data")
		}

		GetUserByEmail, ErrGetUser := auth.UserRepo.FindyByEmail(UserConvert.Credential.Email)

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

	var UserConvert resources.AuthResource

	ValidationErr := service.Validate.Struct(request)

	if ValidationErr != nil {
		return nil, helpers.ValidationError(ctx, ValidationErr), nil
	}

	CheckUserAuthentication, UserErr := service.CredentialRepo.CheckEmailCredential(request.Email)

	if UserErr != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	ErrComparePassword := helpers.ComparePassword(CheckUserAuthentication.Password, request.Password)

	if ErrComparePassword != nil {
		return nil, nil, errors.New("Error uncorrect credential")
	}

	if CheckUserAuthentication.Role.Name == "user" {

		GetUser, _, errUser := service.CredentialRepo.GetAuthUser(CheckUserAuthentication.Id, CheckUserAuthentication.Role.Name)

		if errUser != nil {
			return nil, nil, errors.New("Error uncorrect credential")
		}

		UserConvert = conversionResource.UserDomainToAuthResource(GetUser)

	} else if CheckUserAuthentication.Role.Name == "counselor" {

		_, GetCounselor, errUser := service.CredentialRepo.GetAuthUser(CheckUserAuthentication.Id, CheckUserAuthentication.Role.Name)

		if errUser != nil {
			return nil, nil, errors.New("Error uncorrect credential")
		}

		UserConvert = conversionResource.CounselorDomainToAuthResource(GetCounselor)

	}

	GetToken, ErrToken := helpers.GenerateToken(UserConvert, ctx)

	if ErrToken != nil {
		return nil, nil, errors.New("Failed generate token")
	}

	GetAuthResponse := conversionResource.AuthResourceToAuthTokenResource(UserConvert, GetToken)

	return GetAuthResponse, nil, nil

}
