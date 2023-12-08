package services

import (
	"fmt"
	"mime/multipart"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/pkg/storage"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserService interface {
	RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error)
	GetUserProfile(ctx echo.Context) (*domain.Users, error)
	UpdateUserProfile(ctx echo.Context, request requests.UpdateUserProfileRequest, picture *multipart.FileHeader) (*domain.Users, []exceptions.ValidationMessage, error)
	AddFavoriteArticle(ctx echo.Context, slug string) error
	DeleteFavoriteArticle(ctx echo.Context, slug string) error
	AllFavoriteArticle(ctx echo.Context) (*domain.Users, error)
}

type UserServiceImpl struct {
	UserRepo        repositories.UserRepository
	RoleRepo        repositories.RoleRepository
	Validator       *validator.Validate
	ArticleRepo     repositories.ArticleRepository
	FavoriteArticle repositories.ArticleFavoriteRepository
}

func NewUserService(userServiceImpl UserServiceImpl) UserService {
	return &userServiceImpl
}

func (service *UserServiceImpl) RegisterUser(ctx echo.Context, request requests.UserRequest) (*domain.Users, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingUser, _ := service.UserRepo.FindyByEmail(request.Email)
	if existingUser != nil {
		return nil, nil, fmt.Errorf("Email already exists")
	}

	getRoleUser, _ := service.RoleRepo.FindByName("user")
	if getRoleUser == nil {
		return nil, nil, fmt.Errorf("Role user not found")
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

func (service *UserServiceImpl) UpdateUserProfile(ctx echo.Context, request requests.UpdateUserProfileRequest, picture *multipart.FileHeader) (*domain.Users, []exceptions.ValidationMessage, error) {

	if picture != nil {
		cloudURL, errUpload := storage.S3PutFile(picture, "user/picture")

		if errUpload != nil {
			return nil, nil, errUpload
		}

		request.Profile_picture = cloudURL
	}

	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	getUser, err := service.GetUserProfile(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to find user: %s", err.Error())
	}

	updateProfile := conversion.UserUpdateRequestToUserDomain(request, getUser)

	updateProfile.Credential.Role_id = getUser.Credential.Role_id

	updatedUser, err := service.UserRepo.UpdateUser(updateProfile)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when updating user: %s", err.Error())
	}

	return updatedUser, nil, nil
}

func (service *UserServiceImpl) AddFavoriteArticle(ctx echo.Context, slug string) error {
	getUserClaim := helpers.GetAuthClaims(ctx)

	user, err := service.UserRepo.FindByID(int(getUserClaim.Id))
	if err != nil {
		return fmt.Errorf("Failed to find user: %s", err.Error())
	}

	slugArticle, err := service.ArticleRepo.FindSlugForFavorite(slug)
	if err != nil {
		return fmt.Errorf("Failed to find article: %s", err.Error())
	} else {
		errAddFavorite := service.FavoriteArticle.AddFavoriteArticle(*user, slugArticle)
		if errAddFavorite != nil {
			return errAddFavorite
		}
	}

	return nil
}

func (service *UserServiceImpl) DeleteFavoriteArticle(ctx echo.Context, slug string) error {
	getUserClaim := helpers.GetAuthClaims(ctx)

	user, err := service.UserRepo.FindByID(int(getUserClaim.Id))
	if err != nil {
		return fmt.Errorf("Failed to find user: %s", err.Error())
	}

	slugArticle, err := service.ArticleRepo.FindSlugForFavorite(slug)
	if err != nil {
		return fmt.Errorf("Failed to find article: %s", err.Error())
	} else {
		errAddFavorite := service.FavoriteArticle.DeleteFavoriteArticle(*user, slugArticle)
		if errAddFavorite != nil {
			return errAddFavorite
		}
	}

	return nil
}

func (service *UserServiceImpl) AllFavoriteArticle(ctx echo.Context) (*domain.Users, error) {
	getUserClaim := helpers.GetAuthClaims(ctx)

	user, err := service.UserRepo.FindByID(int(getUserClaim.Id))
	if err != nil {
		return nil, fmt.Errorf("Failed to find user: %s", err.Error())
	}

	return user, nil
}
