package features

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/tests/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserRegister_created(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	validate := validator.New()
	mockRoleRepo := new(mocks.RoleRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  mockUserRepo,
		RoleRepo:  mockRoleRepo,
		Validator: validate,
	})

	mockRoleRepo.On("FindByName", "user").Return(&domain.Roles{
		Id:   1,
		Name: "user",
	}, nil)

	roles, _ := mockRoleRepo.FindByName("user")

	request := requests.UserRequest{
		First_name:   "User",
		Last_name:    "Test",
		Email:        "userTest@gmail.com",
		Username:     "usertest",
		Password:     "userTest123",
		Phone_number: "0812345678890",
		Role_id:      uint(roles.Id),
	}

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	recordHttp := httptest.NewRecorder()

	mockUserRepo.On("FindyByEmail", request.Email).Return(nil, nil)
	mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.Users")).Return(&domain.Users{
		Id:            1,
		Credential_id: 1,
		First_name:    "User",
		Last_name:     "Test",
		Credential: &domain.Credentials{
			Id:       1,
			Email:    "userTest@gmail.com",
			Username: "usertest",
			Password: "userTest123",
			Role_id:  uint(roles.Id),
		},
		Phone_number: "0812345678890",
	}, nil)

	results, _, err := userService.RegisterUser(e.NewContext(req, recordHttp), request)

	assert.NoError(t, err)
	assert.NotNil(t, results)
	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}
func TestUserRegister_error_validation(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	validate := validator.New()
	mockRoleRepo := new(mocks.RoleRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  mockUserRepo,
		RoleRepo:  mockRoleRepo,
		Validator: validate,
	})

	mockRoleRepo.On("FindByName", "user").Return(&domain.Roles{
		Id:   1,
		Name: "user",
	}, nil)

	roles, _ := mockRoleRepo.FindByName("user")

	request := requests.UserRequest{
		First_name:   "",
		Last_name:    "",
		Email:        "userTest@gmail.com",
		Username:     "",
		Password:     "",
		Phone_number: "",
		Role_id:      uint(roles.Id),
	}

	validationErr := validate.Struct(request)
	assert.NotNil(t, validationErr)

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	recordHttp := httptest.NewRecorder()

	results, validation, err := userService.RegisterUser(e.NewContext(req, recordHttp), request)

	assert.NotNil(t, validation)
	assert.NoError(t, err)
	assert.Nil(t, results)

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}

func TestUserRegister_email_exists(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	validate := validator.New()
	mockRoleRepo := new(mocks.RoleRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  mockUserRepo,
		RoleRepo:  mockRoleRepo,
		Validator: validate,
	})

	mockRoleRepo.On("FindByName", "user").Return(&domain.Roles{
		Id:   1,
		Name: "user",
	}, nil)

	roles, _ := mockRoleRepo.FindByName("user")

	request := requests.UserRequest{
		First_name:   "User",
		Last_name:    "Test",
		Email:        "userTest@gmail.com",
		Username:     "usertest",
		Password:     "userTest123",
		Phone_number: "0812345678890",
		Role_id:      uint(roles.Id),
	}

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	recordHttp := httptest.NewRecorder()

	mockUserRepo.On("FindyByEmail", request.Email).Return(&domain.Users{
		Id:            1,
		Credential_id: 1,
		First_name:    "User",
		Last_name:     "Test",
		Credential: &domain.Credentials{
			Id:       1,
			Email:    "userTest@gmail.com",
			Username: "usertest",
			Password: "userTest123",
			Role_id:  uint(roles.Id),
		},
		Phone_number: "0812345678890",
	}, nil)

	results, _, err := userService.RegisterUser(e.NewContext(req, recordHttp), request)

	assert.Error(t, err)
	assert.Nil(t, results)
	assert.EqualError(t, err, "Email already exists")

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}
func TestUserRegister_role_not_found(t *testing.T) {

	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	validate := validator.New()
	mockRoleRepo := new(mocks.RoleRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  mockUserRepo,
		RoleRepo:  mockRoleRepo,
		Validator: validate,
	})

	mockRoleRepo.On("FindByName", "user").Return(nil, fmt.Errorf("Role user not found"))

	_, errRole := mockRoleRepo.FindByName("user")

	assert.NotNil(t, errRole)

	request := requests.UserRequest{
		First_name:   "User",
		Last_name:    "Test",
		Email:        "userTest@gmail.com",
		Username:     "usertest",
		Password:     "userTest123",
		Phone_number: "0812345678890",
	}

	mockUserRepo.On("FindyByEmail", request.Email).Return(nil, nil)

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	recordHttp := httptest.NewRecorder()

	results, _, err := userService.RegisterUser(e.NewContext(req, recordHttp), request)

	assert.Error(t, err)
	assert.Nil(t, results)
	assert.EqualError(t, err, "Role user not found")

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}
func TestUserRegister_failed_created(t *testing.T) {
	e := echo.New()

	mockUserRepo := new(mocks.UserRepository)
	validate := validator.New()
	mockRoleRepo := new(mocks.RoleRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  mockUserRepo,
		RoleRepo:  mockRoleRepo,
		Validator: validate,
	})

	mockRoleRepo.On("FindByName", "user").Return(&domain.Roles{
		Id:   1,
		Name: "user",
	}, nil)

	roles, _ := mockRoleRepo.FindByName("user")

	request := requests.UserRequest{
		First_name:   "User",
		Last_name:    "Test",
		Email:        "userTest@gmail.com",
		Username:     "usertest",
		Password:     "userTest123",
		Phone_number: "0812345678890",
		Role_id:      uint(roles.Id),
	}

	req := httptest.NewRequest(http.MethodPost, "/users/register", nil)
	recordHttp := httptest.NewRecorder()

	mockUserRepo.On("FindyByEmail", request.Email).Return(nil, nil)
	mockUserRepo.On("CreateUser", mock.AnythingOfType("*domain.Users")).Return(nil, fmt.Errorf("ERROR!"))

	_, _, err := userService.RegisterUser(e.NewContext(req, recordHttp), request)

	assert.Error(t, err)

	mockUserRepo.AssertExpectations(t)
	mockRoleRepo.AssertExpectations(t)
}
