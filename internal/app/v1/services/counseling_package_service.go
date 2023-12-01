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

type CounselingPackageService interface {
	CreatePackage(ctx echo.Context, request requests.CounselingPackageRequest, thumbnail *multipart.FileHeader) (*domain.CounselingPackage, []exceptions.ValidationMessage, error)
	GetAllPackage(ctx echo.Context) ([]domain.CounselingPackage, error)
}

type CounselingPackageServiceImpl struct {
	CounselingPackageRepo repositories.CounselingPackageRepository
	Validator             *validator.Validate
	AdminRepo             repositories.AdminRepository
}

func NewCounselingPackageService(counselingpackageServiceImpl CounselingPackageServiceImpl) CounselingPackageService {
	return &counselingpackageServiceImpl
}

func (service *CounselingPackageServiceImpl) CreatePackage(ctx echo.Context, request requests.CounselingPackageRequest, thumbnail *multipart.FileHeader) (*domain.CounselingPackage, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingName, _ := service.CounselingPackageRepo.FindByName(request.Package_name)
	if existingName != nil {
		return nil, nil, fmt.Errorf("Package name already exists")
	}

	ThumbnailCloudURL, errUploadThumbnail := storage.S3PutFile(thumbnail, "package/thumbnail")

	if errUploadThumbnail != nil {
		return nil, nil, errUploadThumbnail
	}

	request.Thumbnail = &ThumbnailCloudURL

	counselingpackage := conversion.PackageCreateRequestToPackageDomain(request)

	result, err := service.CounselingPackageRepo.CreatePackage(counselingpackage)
	if err != nil {
		return nil, nil, fmt.Errorf("Error create counseling package")
	}

	return result, nil, nil
}

func (service *CounselingPackageServiceImpl) GetAllPackage(ctx echo.Context) ([]domain.CounselingPackage, error) {
	result, err := service.CounselingPackageRepo.GetAllPackage()
	if err != nil {
		return nil, err
	}
	return result, nil
}
