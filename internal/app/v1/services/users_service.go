package services

import (
	"fmt"
	"strconv"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error)
	GetUserProfile(ctx echo.Context) (*domain.Users, error)
	UpdateUserProfile(ctx echo.Context, request requests.UpdateUserProfileRequest) (*domain.Users, []exceptions.ValidationMessage, error)
}

type UserServiceImpl struct {
	UserRepo  repositories.UserRepository
	RoleRepo  repositories.RoleRepository
	validator *validator.Validate
}

func NewUserService(user repositories.UserRepository, role repositories.RoleRepository, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepo:  user,
		RoleRepo:  role,
		validator: validator,
	}
}

func (service *UserServiceImpl) RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingUser, _ := service.UserRepo.FindyByEmail(request.Email)
	if existingUser != nil {
		return nil, nil, fmt.Errorf("Email already exists")
	}

	getRoleUser, _ := service.RoleRepo.FindByName("user")
	if getRoleUser == nil {
		return nil, nil, fmt.Errorf("role user not found")
	}

	request.Role_id = uint(getRoleUser.Id)

	user := conversion.UserCreateRequestToUserDomain(request)

	user.Credential.Password = helpers.HashPassword(request.Password)

	result, err := service.UserRepo.CreateUser(user)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when register user: %s", err.Error())
	}

	return result, nil, nil
}

func (s *UserServiceImpl) GetUserProfile(ctx echo.Context) (*domain.Users, error) {
	getUserClaim := helpers.GetAuthClaims(ctx)

	user, err := s.UserRepo.FindByID(int(getUserClaim.Id))
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserServiceImpl) UpdateUserProfile(ctx echo.Context, request requests.UpdateUserProfileRequest) (*domain.Users, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	// Ambil ID dari URL menggunakan ctx.Param("id")
	userID := ctx.Param("id")

	// Lakukan konversi dari string ke tipe data yang sesuai (misalnya, uint)
	parsedUserID, err := strconv.ParseUint(userID, 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("id user not falid")
	}

	existingUser, err := service.UserRepo.FindByID(int(parsedUserID))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to find user: %s", err.Error())
	}

	// Jika pengguna tidak ditemukan, Anda dapat memberikan tanggapan bahwa pengguna tidak ditemukan
	if existingUser == nil {
		return nil, nil, fmt.Errorf("user not found")
	}

	//Perbarui nilai-nilai pengguna yang ada dengan data baru dari request
	existingUser.First_name = request.First_name
	existingUser.Last_name = request.Last_name
	existingUser.Credential.Username = request.Username
	existingUser.Credential.Email = request.Email
	existingUser.Birthday = request.Birthday
	existingUser.Profile_picture = request.Profile_picture
	// Update nilai-nilai lain sesuai kebutuhan

	// Lakukan operasi update ke dalam database
	updatedUser, err := service.UserRepo.UpdateUser(existingUser)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return updatedUser, nil, nil

	// user := conversion.UserUpdateRequestToUserDomain(request)

	// updatedUser, err := service.UserRepo.UpdateUser(user)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("error when update user: %s", err.Error())
	// }

	// return updatedUser, nil, nil
}
