package features

import (
	"fmt"
	"testing"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/tests/mocks"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegisterCounselor_errorValidation(t *testing.T) {

	e := echo.New()

	mockCounseloRepo := new(mocks.CounselorRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	validate := validator.New()

	counselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo: mockCounseloRepo,
		RoleRepo:      mockRoleRepo,
		Validator:     validate,
	})

	mockRoleRepo.On("FindByName", "counselor").Return(&domain.Roles{
		Id:   2,
		Name: "counselor",
	}, nil)

	getRole, _ := mockRoleRepo.FindByName("counselor")

	request := requests.CounselorRequest{
		First_name:  "",
		Last_name:   "",
		Email:       "",
		Description: "",
		Username:    "",
		Password:    "",
		Role_id:     uint(getRole.Id),
	}

	validationErr := validate.Struct(request)
	assert.NotNil(t, validationErr)

	_, validation, err := counselorService.RegisterCounselor(e.NewContext(nil, nil), request)

	assert.NoError(t, err)
	assert.NotNil(t, validation)
	mockRoleRepo.AssertExpectations(t)
	mockCounseloRepo.AssertExpectations(t)
}
func TestRegisterCounselor_emailAlreadyExists(t *testing.T) {

	e := echo.New()

	mockCounseloRepo := new(mocks.CounselorRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	validate := validator.New()

	counselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo: mockCounseloRepo,
		RoleRepo:      mockRoleRepo,
		Validator:     validate,
	})

	mockRoleRepo.On("FindByName", "counselor").Return(&domain.Roles{
		Id:   2,
		Name: "counselor",
	}, nil)

	getRole, _ := mockRoleRepo.FindByName("counselor")

	request := requests.CounselorRequest{
		First_name:  "Counselor",
		Last_name:   "Test 1",
		Email:       "counselor@gmail.com",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Username:    "counselortest1",
		Password:    "counselortest1123",
		Role_id:     uint(getRole.Id),
	}

	HashPassword := helpers.HashPassword(request.Password)

	request.Password = HashPassword

	mockCounseloRepo.On("FindyByEmail", request.Email).Return(&domain.Counselors{
		Id:            1,
		Credential_id: 2,
		First_name:    "Counselor",
		Last_name:     "Test 2",
		Description:   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Credential: &domain.Credentials{
			Id:       2,
			Username: "counselortest1",
			Password: HashPassword,
			Role_id:  uint(getRole.Id),
		},
		Education: "Universitas Udayana",
	}, nil)

	_, _, err := counselorService.RegisterCounselor(e.NewContext(nil, nil), request)

	assert.Error(t, err)

	mockRoleRepo.AssertExpectations(t)
	mockCounseloRepo.AssertExpectations(t)
}
func TestRegisterCounselor_RoleNotFound(t *testing.T) {

	e := echo.New()

	mockCounseloRepo := new(mocks.CounselorRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	validate := validator.New()

	counselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo: mockCounseloRepo,
		RoleRepo:      mockRoleRepo,
		Validator:     validate,
	})

	mockRoleRepo.On("FindByName", "counselor").Return(nil, fmt.Errorf("Roles not found"))

	getRole, errRoles := mockRoleRepo.FindByName("counselor")

	assert.Error(t, errRoles)
	assert.Nil(t, getRole)

	request := requests.CounselorRequest{
		First_name:  "Counselor",
		Last_name:   "Test 1",
		Email:       "counselor@gmail.com",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Username:    "counselortest1",
		Password:    "counselortest1123",
	}

	HashPassword := helpers.HashPassword(request.Password)

	request.Password = HashPassword

	mockCounseloRepo.On("FindyByEmail", request.Email).Return(nil, nil)

	_, _, err := counselorService.RegisterCounselor(e.NewContext(nil, nil), request)

	assert.Error(t, err)

	mockRoleRepo.AssertExpectations(t)
	mockCounseloRepo.AssertExpectations(t)
}
func TestRegisterCounselor_created(t *testing.T) {

	e := echo.New()

	mockCounseloRepo := new(mocks.CounselorRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	validate := validator.New()

	counselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo: mockCounseloRepo,
		RoleRepo:      mockRoleRepo,
		Validator:     validate,
	})

	mockRoleRepo.On("FindByName", "counselor").Return(&domain.Roles{
		Id:   2,
		Name: "counselor",
	}, nil)

	getRole, _ := mockRoleRepo.FindByName("counselor")

	request := requests.CounselorRequest{
		First_name:  "Counselor",
		Last_name:   "Test 1",
		Email:       "counselor@gmail.com",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Username:    "counselortest1",
		Password:    "counselortest1123",
		Role_id:     uint(getRole.Id),
	}

	HashPassword := helpers.HashPassword(request.Password)

	request.Password = HashPassword

	mockCounseloRepo.On("FindyByEmail", request.Email).Return(nil, nil)
	mockCounseloRepo.On("CreateCounselor", mock.AnythingOfType("*domain.Counselors")).Return(&domain.Counselors{
		Id:            1,
		Credential_id: 2,
		First_name:    "Counselor",
		Last_name:     "Test 2",
		Description:   "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Credential: &domain.Credentials{
			Id:       2,
			Username: "counselortest1",
			Password: HashPassword,
			Role_id:  uint(getRole.Id),
		},
		Education: "Universitas Udayana",
	}, nil)

	result, _, err := counselorService.RegisterCounselor(e.NewContext(nil, nil), request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	mockRoleRepo.AssertExpectations(t)
	mockCounseloRepo.AssertExpectations(t)
}
func TestRegisterCounselor_errorCreated(t *testing.T) {

	e := echo.New()

	mockCounseloRepo := new(mocks.CounselorRepository)
	mockRoleRepo := new(mocks.RoleRepository)
	validate := validator.New()

	counselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo: mockCounseloRepo,
		RoleRepo:      mockRoleRepo,
		Validator:     validate,
	})

	mockRoleRepo.On("FindByName", "counselor").Return(&domain.Roles{
		Id:   2,
		Name: "counselor",
	}, nil)

	getRole, _ := mockRoleRepo.FindByName("counselor")

	request := requests.CounselorRequest{
		First_name:  "Counselor",
		Last_name:   "Test 1",
		Email:       "counselor@gmail.com",
		Description: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book.",
		Username:    "counselortest1",
		Password:    "counselortest1123",
		Role_id:     uint(getRole.Id),
	}

	HashPassword := helpers.HashPassword(request.Password)

	request.Password = HashPassword

	mockCounseloRepo.On("FindyByEmail", request.Email).Return(nil, nil)
	mockCounseloRepo.On("CreateCounselor", mock.AnythingOfType("*domain.Counselors")).Return(nil, fmt.Errorf("ERROR"))

	result, _, err := counselorService.RegisterCounselor(e.NewContext(nil, nil), request)

	assert.Error(t, err)
	assert.Nil(t, result)

	mockRoleRepo.AssertExpectations(t)
	mockCounseloRepo.AssertExpectations(t)
}
