package services

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"strings"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	conRes "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/pkg/storage"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CounselorService interface {
	RegisterCounselor(ctx echo.Context, request requests.CounselorRequest) (*domain.Counselors, []exceptions.ValidationMessage, error)
	AddSpecialist(ctx echo.Context, id uint, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error)
	RemoveSpecialistCounselor(ctx echo.Context, id int, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error)
	GetAllCounselors(ctx echo.Context) ([]domain.Counselors, error)
	GetCounselorsForMobile(ctx echo.Context) ([]domain.Counselors, error)
	GetCounselorProfile(ctx echo.Context) (*domain.Counselors, error)
	GetDetailCounselor(ctx echo.Context) (*resources.DetailCounselor, error)
	GetDetailCounselorOnly(ctx echo.Context) (*resources.DetailCounselor, error)
	UpdateCounselor(ctx echo.Context, request requests.UpdateCounselorProfileRequest, picture *multipart.FileHeader) (*domain.Counselors, []exceptions.ValidationMessage, error)
	UpdateCounselorForMobile(ctx echo.Context, request requests.UpdateCounselorProfileRequestForMobile, picture *multipart.FileHeader) (*domain.Counselors, []exceptions.ValidationMessage, error)
}

type CounselorServiceImpl struct {
	ScheduleRepo               repositories.ScheduleRepository
	CounselorRepo              repositories.CounselorRepository
	RoleRepo                   repositories.RoleRepository
	Validator                  *validator.Validate
	AdminRepo                  repositories.AdminRepository
	SpecialistRepo             repositories.SpecialistRepository
	CounselorHasSpecialistRepo repositories.CounseloHasSpecialistRepository
}

func NewCounselorService(counselorServiceImpl CounselorServiceImpl) CounselorService {
	return &counselorServiceImpl
}

func (service *CounselorServiceImpl) GetAllCounselors(ctx echo.Context) ([]domain.Counselors, error) {
	FilterSpecialist := requests.FilterCounselorsSpecialist{}

	Specialist := ctx.QueryParam("specialist")
	Search := ctx.QueryParam("search")

	if Specialist != "" {
		FilterSpecialist.Specialist = strings.Split(Specialist, ",")
	}

	counselors, err := service.CounselorRepo.FindAllCounselors(Search, FilterSpecialist.Specialist)
	if err != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	return counselors, nil
}

func (service *CounselorServiceImpl) RemoveSpecialistCounselor(ctx echo.Context, id int, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error) {
	ValidationMessage := service.Validator.Struct(request)
	var Specialist []domain.Specialist

	if ValidationMessage != nil || len(request.Name) == 0 {
		return helpers.ValidationError(ctx, ValidationMessage), nil
	}

	GetCounselorData, errGetCounselor := service.CounselorRepo.FindById(id)

	if errGetCounselor != nil {
		fmt.Errorf(errGetCounselor.Error())
		return nil, fmt.Errorf("Counselor not found")
	}

	for _, val := range GetCounselorData.Specialists {

		for index := range request.Name {

			if request.Name[index] == val.Name {

				GetSpecialists, errGetSpecialists := service.SpecialistRepo.FindSpecialistByName(val.Name)

				if errGetSpecialists != nil {
					fmt.Println(errGetSpecialists.Error())
					return nil, fmt.Errorf("One of Counselor request is not found")
				}

				Specialist = append(Specialist, *GetSpecialists)

			}
		}

	}

	if len(Specialist) <= 0 {
		return nil, fmt.Errorf("One of counselor request is not found")
	}

	ErrRemoveCounselor := service.CounselorHasSpecialistRepo.RemoveManySpecialist(*GetCounselorData, Specialist)

	if ErrRemoveCounselor != nil {
		fmt.Errorf(ErrRemoveCounselor.Error())
		return nil, fmt.Errorf("Error when remove specialist")
	}

	return nil, nil
}

func (service *CounselorServiceImpl) RegisterCounselor(ctx echo.Context, request requests.CounselorRequest) (*domain.Counselors, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingCounselor, _ := service.CounselorRepo.FindyByEmail(request.Email)
	if existingCounselor != nil {
		return nil, nil, fmt.Errorf("Email already exist")
	}

	getRoleUser, _ := service.RoleRepo.FindByName("counselor")
	if getRoleUser == nil {
		return nil, nil, fmt.Errorf("Role user not found")
	}

	request.Role_id = uint(getRoleUser.Id)

	counselor := conversion.CounselorCreateRequestToCounselorDomain(request)

	counselor.Credential.Password = helpers.HashPassword(request.Password)

	result, err := service.CounselorRepo.CreateCounselor(counselor)
	if err != nil {
		return nil, nil, fmt.Errorf("Error when register counselor: %s", err.Error())
	}

	return result, nil, nil
}

func (service *CounselorServiceImpl) AddSpecialist(ctx echo.Context, id uint, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil || len(request.Name) == 0 {
		return helpers.ValidationError(ctx, err), nil
	}

	counselor, errCounselor := service.CounselorRepo.FindById(int(id))
	if errCounselor != nil {
		return nil, errCounselor
	}

	for index := range request.Name {
		GetSpecialists, errGetSpecialists := service.SpecialistRepo.FindSpecialistByName(request.Name[index])
		if errGetSpecialists == nil {
			errAdd := service.CounselorHasSpecialistRepo.AddSpecialist(*counselor, GetSpecialists)
			if errAdd != nil {
				return nil, errAdd
			}
		} else {
			fmt.Println(errGetSpecialists.Error())
			return nil, fmt.Errorf("One of Counselor request is not found")
		}
	}
	return nil, nil
}

func (service *CounselorServiceImpl) GetCounselorProfile(ctx echo.Context) (*domain.Counselors, error) {
	getUserClaim := helpers.GetAuthClaims(ctx)

	counselor, err := service.CounselorRepo.FindById(int(getUserClaim.Id))
	if err != nil {
		return nil, err
	}
	return counselor, nil
}

func (service *CounselorServiceImpl) UpdateCounselor(ctx echo.Context, request requests.UpdateCounselorProfileRequest, picture *multipart.FileHeader) (*domain.Counselors, []exceptions.ValidationMessage, error) {
	if picture != nil {
		cloudURL, errUpload := storage.S3PutFile(picture, "counselor/picture")

		if errUpload != nil {
			return nil, nil, fmt.Errorf("Error when upload picture: %s", errUpload.Error())
		}

		request.Profile_picture = cloudURL
	}

	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	getId := ctx.Param("id")
	getcounselorId, _ := strconv.Atoi(getId)

	getUser, _ := service.CounselorRepo.FindById(getcounselorId)
	if getUser == nil {
		return nil, nil, fmt.Errorf("Counselor not found")
	}
	request.Role_id = getUser.Credential.Role_id

	encryptPassword := helpers.HashPassword(request.Password)
	request.Password = encryptPassword

	counselor := conversion.CounselorUpdateRequestToCounselorDomain(request, getUser)

	errUpdate := service.CounselorRepo.UpdateCounselor(counselor)
	if errUpdate != nil {
		return nil, nil, fmt.Errorf("Error when update counselor: %s", errUpdate.Error())
	}

	return nil, nil, nil
}

func (service *CounselorServiceImpl) UpdateCounselorForMobile(ctx echo.Context, request requests.UpdateCounselorProfileRequestForMobile, picture *multipart.FileHeader) (*domain.Counselors, []exceptions.ValidationMessage, error) {
	if picture != nil {
		cloudURL, errUpload := storage.S3PutFile(picture, "counselor/picture")

		if errUpload != nil {
			return nil, nil, fmt.Errorf("Error when upload picture: %s", errUpload.Error())
		}

		request.Profile_picture = cloudURL
	}

	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	getUser, err := service.GetCounselorProfile(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("Failed to find counselor: %s", err.Error())
	}

	request.Role_id = getUser.Credential.Role_id
	counselor := conversion.CounselorUpdateRequestToCounselorDomainForMobile(request, getUser)

	errUpdate := service.CounselorRepo.UpdateCounselor(counselor)

	if errUpdate != nil {
		return nil, nil, fmt.Errorf("Error when update counselor: %s", errUpdate.Error())
	}

	return nil, nil, nil
}

func (service *CounselorServiceImpl) GetCounselorsForMobile(ctx echo.Context) ([]domain.Counselors, error) {
	FilterSpecialist := requests.FilterCounselorsSpecialist{}

	Specialist := ctx.QueryParam("specialist")
	Search := ctx.QueryParam("search")

	if Specialist != "" {
		FilterSpecialist.Specialist = strings.Split(Specialist, ",")
	}

	counselors, err := service.CounselorRepo.FindAllCounselors(Search, FilterSpecialist.Specialist)
	if err != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	return counselors, nil
}

func (service *CounselorServiceImpl) GetDetailCounselor(ctx echo.Context) (*resources.DetailCounselor, error) {
	getId := ctx.Param("id")
	getcounselorId, _ := strconv.Atoi(getId)

	getCounselor, _ := service.CounselorRepo.FindById(getcounselorId)
	if getCounselor == nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	getScheduleCounselor, errGetSchedule := service.ScheduleRepo.GroupingStartTimeAndFinishTimeCounseling(getcounselorId)

	if errGetSchedule != nil {
		return nil, errGetSchedule
	}

	counselorResponse := conRes.ConvertCounselorDomainToCounselorDetailResponse(getCounselor, getScheduleCounselor)

	return &counselorResponse, nil
}

func (service *CounselorServiceImpl) GetDetailCounselorOnly(ctx echo.Context) (*resources.DetailCounselor, error) {

	getId := ctx.Param("id")
	getcounselorId, _ := strconv.Atoi(getId)

	getCounselor, _ := service.CounselorRepo.FindById(getcounselorId)
	if getCounselor == nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	counselorResponse := conRes.ConvertCounselorDomainToDetail(getCounselor)

	return &counselorResponse, nil

}
