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

type CareerService interface {
	CreateCareer(ctx echo.Context, request requests.CareerRequest, logo *multipart.FileHeader, cover *multipart.FileHeader) (*domain.Career, []exceptions.ValidationMessage, error)
	FindAllCareer(ctx echo.Context) ([]domain.Career, error)
}

type CareerServiceImpl struct {
	CareerRepo repositories.CareerRepository
	Validator  *validator.Validate
}

func NewCareerService(careerServiceImpl CareerServiceImpl) CareerService {
	return &careerServiceImpl
}

func (service *CareerServiceImpl) CreateCareer(ctx echo.Context, request requests.CareerRequest, logo *multipart.FileHeader, cover *multipart.FileHeader) (*domain.Career, []exceptions.ValidationMessage, error) {
	LogoCloudURL, errUploadLogo := storage.
		S3PutFile(logo, "career/logo")

	if errUploadLogo != nil {
		return nil, nil, errUploadLogo
	}

	request.Logo = &LogoCloudURL

	CoverCloudURL, errUploadCover := storage.S3PutFile(cover, "career/cover")

	if errUploadCover != nil {
		return nil, nil, errUploadCover
	}

	request.Cover = &CoverCloudURL
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	career := conversion.CareerCreateRequestToCareerDomain(request)

	createCareer, err := service.CareerRepo.CreateCareer(career)

	if err != nil {
		return nil, nil, fmt.Errorf("error create career: %w", err)
	}

	return createCareer, nil, nil
}

func (service *CareerServiceImpl) FindAllCareer(ctx echo.Context) ([]domain.Career, error) {

	career, err := service.CareerRepo.GetAllCareer()

	if err != nil {
		return nil, fmt.Errorf("error get all career: %w", err)
	}

	return career, nil
}
