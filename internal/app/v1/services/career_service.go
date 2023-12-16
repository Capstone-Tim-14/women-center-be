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

type CareerService interface {
	CreateCareer(ctx echo.Context, request requests.CareerRequest, logo *multipart.FileHeader, cover *multipart.FileHeader) (*domain.Career, []exceptions.ValidationMessage, error)
	FindAllCareer(ctx echo.Context) ([]domain.Career, error)
	FindCareerByid(ctx echo.Context, id int) (*domain.Career, error)
	AddJobType(ctx echo.Context, id int, request requests.CareerhasTypeRequest) ([]exceptions.ValidationMessage, error)
	RemoveJobType(ctx echo.Context, id int, request requests.CareerhasManyRequest) ([]exceptions.ValidationMessage, error)
	UpdateCareer(ctx echo.Context, request requests.CareerRequest, logo *multipart.FileHeader, cover *multipart.FileHeader) ([]exceptions.ValidationMessage, error)
	DeleteCareer(ctx echo.Context) error
	RecomendationCareerList(ctx echo.Context) ([]domain.Career, error)
	UpdateRecomendationCareer(ctx echo.Context) error
	RecomendationCareerListForMobile(ctx echo.Context) ([]domain.Career, error)
}

type CareerServiceImpl struct {
	CareerRepo        repositories.CareerRepository
	CareerhasTypeRepo repositories.CareerhasTypeRepository
	JobTypeRepo       repositories.JobTypeRepository
	Validator         *validator.Validate
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
		return nil, nil, fmt.Errorf("Error create career: %w", err)
	}

	return createCareer, nil, nil
}

func (service *CareerServiceImpl) FindAllCareer(ctx echo.Context) ([]domain.Career, error) {

	var FilterCareer requests.CareerFilterRequest

	JobType := ctx.QueryParam("job_type")

	if JobType != "" {
		FilterCareer.JobType = strings.Split(JobType, ",")
	}

	career, err := service.CareerRepo.GetAllCareer(FilterCareer)

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, fmt.Errorf("Career is empty")
	}

	return career, nil
}

func (service *CareerServiceImpl) FindCareerByid(ctx echo.Context, id int) (*domain.Career, error) {

	careerDetail, err := service.CareerRepo.FindCareerByid(id)

	if err != nil {
		return nil, fmt.Errorf("Error get detail career: %w", err)
	}

	return careerDetail, nil
}

func (service *CareerServiceImpl) AddJobType(ctx echo.Context, id int, request requests.CareerhasTypeRequest) ([]exceptions.ValidationMessage, error) {

	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	career, errCareer := service.CareerRepo.FindCareerByid(id)
	if errCareer != nil {
		return nil, errCareer
	}

	jobtype, errJobtype := service.JobTypeRepo.FindJobTypeByName(request.Name)
	if errJobtype != nil {
		return nil, errJobtype
	}

	errAddJobtype := service.CareerhasTypeRepo.AddJobType(*career, jobtype)
	if errAddJobtype != nil {
		return nil, errAddJobtype
	}
	return nil, nil
}

func (service *CareerServiceImpl) RemoveJobType(ctx echo.Context, id int, request requests.CareerhasManyRequest) ([]exceptions.ValidationMessage, error) {

	ValidationMessage := service.Validator.Struct(request)
	var JobTypeList []domain.Job_Type

	if ValidationMessage != nil {
		return helpers.ValidationError(ctx, ValidationMessage), nil
	}

	GetCareer, errGetCareer := service.CareerRepo.FindCareerByid(id)

	if errGetCareer != nil {
		fmt.Errorf(errGetCareer.Error())
		return nil, fmt.Errorf("Career not found")
	}

	for _, val := range GetCareer.Job_type {

		for index := range request.Name {

			if request.Name[index] == val.Name {

				GetJobType, errGetJobType := service.JobTypeRepo.FindJobTypeByName(val.Name)

				if errGetJobType != nil {
					fmt.Println(errGetJobType.Error())
					return nil, fmt.Errorf("One of career request is not found")

				}

				JobTypeList = append(JobTypeList, *GetJobType)

			} else {
				continue
			}
		}
	}

	if len(JobTypeList) <= 0 {
		return nil, fmt.Errorf("One of career request is not found")
	}

	ErrRemoveCareer := service.CareerhasTypeRepo.RemoveJobTypeById(*GetCareer, JobTypeList)

	if ErrRemoveCareer != nil {
		fmt.Errorf(ErrRemoveCareer.Error())
		return nil, fmt.Errorf("Error when remove job type")
	}

	return nil, nil

}

func (service *CareerServiceImpl) UpdateCareer(ctx echo.Context, request requests.CareerRequest, logo *multipart.FileHeader, cover *multipart.FileHeader) ([]exceptions.ValidationMessage, error) {

	if logo != nil {
		LogoCloudURL, errUploadLogo := storage.S3PutFile(logo, "career/logo")

		if errUploadLogo != nil {
			return nil, errUploadLogo
		}

		request.Logo = &LogoCloudURL
	}

	if cover != nil {
		CoverCloudURL, errUploadCover := storage.S3PutFile(cover, "career/cover")

		if errUploadCover != nil {
			return nil, errUploadCover
		}

		request.Cover = &CoverCloudURL
	}

	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	getId := ctx.Param("id")
	updateId, _ := strconv.Atoi(getId)

	_, errCareer := service.CareerRepo.FindCareerByid(updateId)

	if errCareer != nil {
		return nil, fmt.Errorf("Error get detail career: %w", errCareer)
	}

	career := conversion.CareerCreateRequestToCareerDomain(request)

	errUpdateCareer := service.CareerRepo.UpdateCareerById(updateId, career)

	if errUpdateCareer != nil {
		return nil, fmt.Errorf("Error update career: %w", errUpdateCareer)
	}

	return nil, nil
}

func (service *CareerServiceImpl) DeleteCareer(ctx echo.Context) error {

	getId := ctx.Param("id")
	deleteId, _ := strconv.Atoi(getId)

	_, errCareer := service.CareerRepo.FindCareerByid(deleteId)

	if errCareer != nil {
		return fmt.Errorf("Error get detail career: %w", errCareer)
	}

	errDeleteCareer := service.CareerRepo.DeleteCareerById(deleteId)

	if errDeleteCareer != nil {
		return fmt.Errorf("Error delete career: %w", errDeleteCareer)
	}

	return nil
}

func (service *CareerServiceImpl) RecomendationCareerList(ctx echo.Context) ([]domain.Career, error) {

	var FilterCareer requests.CareerFilterRequest

	JobType := ctx.QueryParam("job_type")

	if JobType != "" {
		FilterCareer.JobType = strings.Split(JobType, ",")
	}

	career, err := service.CareerRepo.RecomendationCareerList(FilterCareer)

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, fmt.Errorf("Career is empty")
	}

	return career, nil
}

func (service *CareerServiceImpl) UpdateRecomendationCareer(ctx echo.Context) error {

	getId := ctx.QueryParam("career_id")
	updateId, _ := strconv.Atoi(getId)

	career, errCareer := service.CareerRepo.FindCareerByid(updateId)

	if errCareer != nil {
		return fmt.Errorf("Error get detail career: %w", errCareer)
	}
	getStatus := ctx.QueryParam("status")
	updateStatus, _ := strconv.ParseBool(getStatus)

	errUpdateCareer := service.CareerRepo.UpdateRecomendationCareer(updateStatus, career)

	if errUpdateCareer != nil {
		return fmt.Errorf("Error update career: %w", errUpdateCareer)
	}

	return nil
}

func (service *CareerServiceImpl) RecomendationCareerListForMobile(ctx echo.Context) ([]domain.Career, error) {

	var FilterCareer requests.CareerFilterRequest

	JobType := ctx.QueryParam("job_type")

	if JobType != "" {
		FilterCareer.JobType = strings.Split(JobType, ",")
	}

	career, err := service.CareerRepo.RecomendationCareerList(FilterCareer)

	if err != nil {
		fmt.Errorf(err.Error())
		return nil, fmt.Errorf("Career is empty")
	}

	return career, nil
}
