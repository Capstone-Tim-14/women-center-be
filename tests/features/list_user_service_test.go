package features

import (
	"fmt"
	"testing"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/tests/mocks"

	"github.com/stretchr/testify/assert"
)

func TestGetAllUser_Success(t *testing.T) {

	// Create mock repository
	mockUserRepo := new(mocks.UserRepository)

	// Create new service and set parameter to mock repo
	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo: mockUserRepo,
	})

	// Setup return in mock method inside
	mockUserRepo.On("GetAllUsers").Return([]domain.Users{
		{
			Id:            1,
			Credential_id: 1,
			Credential: &domain.Credentials{
				Id:       1,
				Username: "test1",
				Email:    "test1@gmail.com",
			},
			First_name:   "Test",
			Last_name:    "User",
			Phone_number: "08123456789",
		},
	}, nil)

	// call service
	results, err := userService.UsersList()

	// Assert the results
	assert.NoError(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, len(results))

	mockUserRepo.AssertExpectations(t)

}
func TestGetAllUser_error(t *testing.T) {

	mockUserRepo := new(mocks.UserRepository)

	userService := services.NewUserService(services.UserServiceImpl{
		UserRepo: mockUserRepo,
	})

	mockUserRepo.On("GetAllUsers").Return(nil, fmt.Errorf("Error when get all user"))

	results, err := userService.UsersList()

	assert.Error(t, err)
	assert.Nil(t, results)
	assert.EqualError(t, err, "Error when get all user")

	mockUserRepo.AssertExpectations(t)

}
