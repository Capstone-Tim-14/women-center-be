package features

import (
	"fmt"
	"testing"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/tests/mocks"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserAuthentication_ValidationError(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "",
		Password: "",
	}

	validation := validate.Struct(request)

	assert.Error(t, validation)

	resultToken, errVal, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.Nil(t, resultToken)
	assert.NotNil(t, errVal)
	assert.NoError(t, err)

	mockCredentialRepo.AssertExpectations(t)
}
func TestUserAuthentication_uncorrectEmail(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "usertest1@gmail.com",
		Password: "usertest123",
	}

	mockCredentialRepo.On("CheckEmailCredential", request.Email).Return(nil, fmt.Errorf("User does not exists by email"))

	resultToken, _, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("Error uncorrect credential"))
	assert.Nil(t, resultToken)

	mockCredentialRepo.AssertExpectations(t)
}
func TestUserAuthentication_uncorrectPassword(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "usertest1@gmail.com",
		Password: "usertest123",
	}

	hashPassword := helpers.HashPassword("userTest123")

	mockCredentialRepo.On("CheckEmailCredential", request.Email).Return(&domain.Credentials{
		Id:       1,
		Username: "usertest1",
		Email:    "usertest1@gmail.com",
		Role_id:  1,
		Password: hashPassword,
		Role: &domain.Roles{
			Id:   1,
			Name: "user",
		},
	}, nil)

	errComparePassword := helpers.ComparePassword(hashPassword, request.Password)
	assert.Error(t, errComparePassword)

	resultToken, _, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("Error uncorrect credential"))
	assert.Nil(t, resultToken)

	mockCredentialRepo.AssertExpectations(t)
}
func TestUserAuthentication_ErrorGetAuthUser(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "usertest1@gmail.com",
		Password: "usertest123",
	}

	hashPassword := helpers.HashPassword("usertest123")

	mockCredentialRepo.On("CheckEmailCredential", request.Email).Return(&domain.Credentials{
		Id:       1,
		Username: "usertest1",
		Email:    "usertest1@gmail.com",
		Role_id:  1,
		Password: hashPassword,
		Role: &domain.Roles{
			Id:   1,
			Name: "user",
		},
	}, nil)

	getUser, _ := mockCredentialRepo.CheckEmailCredential(request.Email)
	assert.NotNil(t, getUser)

	errComparePassword := helpers.ComparePassword(hashPassword, request.Password)
	assert.NoError(t, errComparePassword)

	mockCredentialRepo.On("GetAuthUser", mock.AnythingOfType("uint"), mock.AnythingOfType("string")).Return(nil, nil, fmt.Errorf("Error user not found"))

	_, _, errGetUser := mockCredentialRepo.GetAuthUser(uint(2), "user")

	assert.Error(t, errGetUser)

	_, _, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.Error(t, err)
	assert.Equal(t, err, fmt.Errorf("Error uncorrect credential"))

	mockCredentialRepo.AssertExpectations(t)
}
func TestUserAuthentication_successAuthAsUser(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "usertest1@gmail.com",
		Password: "usertest123",
	}

	hashPassword := helpers.HashPassword("usertest123")

	mockCredentialRepo.On("CheckEmailCredential", request.Email).Return(&domain.Credentials{
		Id:       1,
		Username: "usertest1",
		Email:    "usertest1@gmail.com",
		Role_id:  1,
		Password: hashPassword,
		Role: &domain.Roles{
			Id:   1,
			Name: "user",
		},
	}, nil)

	getUser, _ := mockCredentialRepo.CheckEmailCredential(request.Email)

	assert.NotNil(t, getUser)

	errComparePassword := helpers.ComparePassword(hashPassword, request.Password)
	assert.NoError(t, errComparePassword)

	mockCredentialRepo.On("GetAuthUser", getUser.Id, getUser.Role.Name).Return(&domain.Users{
		Id:            1,
		Credential_id: 1,
		First_name:    "User",
		Last_name:     "Test1",
		Credential: &domain.Credentials{
			Id:       1,
			Username: "usertest1",
			Email:    "usertest1@gmail.com",
			Role_id:  1,
			Password: hashPassword,
			Role: &domain.Roles{
				Id:   1,
				Name: "user",
			},
		},
	}, nil, nil)

	getAuthUser, _, err := mockCredentialRepo.GetAuthUser(getUser.Id, getUser.Role.Name)

	assert.NoError(t, err)
	assert.NotNil(t, getAuthUser)

	convertAuth := conversion.UserDomainToAuthResource(getAuthUser)

	getToken, errToken := helpers.GenerateToken(convertAuth, e.NewContext(nil, nil))

	assert.NoError(t, errToken)
	assert.NotNil(t, getToken)

	resultToken, _, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.NoError(t, err)
	assert.NotNil(t, resultToken)

	mockCredentialRepo.AssertExpectations(t)
}
func TestUserAuthentication_successAuthAsCounselor(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	mockCredentialRepo := new(mocks.CredentialRepository)

	validate := validator.New()

	authService := services.NewAuthService(services.AuthServiceImpl{
		UserRepo:       mockUserRepo,
		RoleRepo:       mockRoleRepo,
		Validate:       validate,
		CredentialRepo: mockCredentialRepo,
	})

	request := requests.AuthRequest{
		Email:    "usertest2@gmail.com",
		Password: "usertest123",
	}

	hashPassword := helpers.HashPassword("usertest123")

	mockCredentialRepo.On("CheckEmailCredential", request.Email).Return(&domain.Credentials{
		Id:       2,
		Username: "usertest2",
		Email:    "usertest2@gmail.com",
		Role_id:  2,
		Password: hashPassword,
		Role: &domain.Roles{
			Id:   2,
			Name: "counselor",
		},
	}, nil)

	getUser, _ := mockCredentialRepo.CheckEmailCredential(request.Email)

	assert.NotNil(t, getUser)

	errComparePassword := helpers.ComparePassword(hashPassword, request.Password)
	assert.NoError(t, errComparePassword)

	mockCredentialRepo.On("GetAuthUser", getUser.Id, getUser.Role.Name).Return(nil, &domain.Counselors{
		Id:            2,
		Credential_id: 2,
		First_name:    "User",
		Last_name:     "Test1",
		Credential: &domain.Credentials{
			Id:       2,
			Username: "usertest1",
			Email:    "usertest1@gmail.com",
			Role_id:  2,
			Password: hashPassword,
			Role: &domain.Roles{
				Id:   2,
				Name: "user",
			},
		},
	}, nil)

	_, getAuthCounselor, err := mockCredentialRepo.GetAuthUser(getUser.Id, getUser.Role.Name)

	assert.NoError(t, err)
	assert.NotNil(t, getAuthCounselor)

	convertAuth := conversion.CounselorDomainToAuthResource(getAuthCounselor)

	getToken, errToken := helpers.GenerateToken(convertAuth, e.NewContext(nil, nil))

	assert.NoError(t, errToken)
	assert.NotNil(t, getToken)

	resultToken, _, err := authService.UserAuthentication(request, e.NewContext(nil, nil))

	assert.NoError(t, err)
	assert.NotNil(t, resultToken)

	mockCredentialRepo.AssertExpectations(t)
}
