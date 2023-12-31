package services

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
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
	FindByTitle(ctx echo.Context, title string) ([]domain.CounselingPackage, error)
	GetAllPackage(ctx echo.Context) ([]domain.CounselingPackage, error)
	GetPackageById(ctx echo.Context, id int) (*domain.CounselingPackage, error)
	DeletePackageById(ctx echo.Context, id int) error
	UpdatePackageById(ctx echo.Context, request requests.CounselingPackageRequest, thumbnail *multipart.FileHeader) (*domain.CounselingPackage, []exceptions.ValidationMessage, error)
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

	existingName, _ := service.CounselingPackageRepo.FindByTitle(request.Title)
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

func (service *CounselingPackageServiceImpl) FindByTitle(ctx echo.Context, title string) ([]domain.CounselingPackage, error) {
	result, err := service.CounselingPackageRepo.FindByTitle(title)
	if err != nil {
		return nil, fmt.Errorf("Package not found")
	}

	return result, nil
}

func (service *CounselingPackageServiceImpl) GetAllPackage(ctx echo.Context) ([]domain.CounselingPackage, error) {
	result, err := service.CounselingPackageRepo.GetAllPackage()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (service *CounselingPackageServiceImpl) DeletePackageById(ctx echo.Context, id int) error {

	err := service.CounselingPackageRepo.DeletePackageById(id)
	if err != nil {
		if strings.Contains(err.Error(), "package not found") {
			return fmt.Errorf("package not found")
		}
		return fmt.Errorf("error when delete package")
	}

	return nil
}

func (service *CounselingPackageServiceImpl) GetPackageById(ctx echo.Context, id int) (*domain.CounselingPackage, error) {
	idPackage, err := service.CounselingPackageRepo.FindById(id)
	if err != nil {
		return nil, err
	}
	return idPackage, nil
}

func (service *CounselingPackageServiceImpl) UpdatePackageById(ctx echo.Context, request requests.CounselingPackageRequest, thumbnail *multipart.FileHeader) (*domain.CounselingPackage, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	id := ctx.Param("id")
	getId, errId := strconv.Atoi(id)
	if errId != nil {
		return nil, nil, fmt.Errorf("invalid id")
	}

	_, err = service.CounselingPackageRepo.FindById(getId)
	if err != nil {
		return nil, nil, fmt.Errorf("package not found")
	}

	if thumbnail != nil {

		ThumbnailCloudURL, errUploadThumbnail := storage.S3PutFile(thumbnail, "counseling_packages/thumbnail")

		if errUploadThumbnail != nil {
			return nil, nil, errUploadThumbnail
		}

		request.Thumbnail = &ThumbnailCloudURL
	}
	cpackage := conversion.PackageCreateRequestToPackageDomain(request)

	_, err = service.CounselingPackageRepo.UpdatePackage(getId, cpackage), nil
	if err != nil {
		return nil, nil, fmt.Errorf("Error update package")
	}

	return nil, nil, nil
}
