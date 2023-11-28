package services

import (
	"errors"
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type CounselorService interface {
	RegisterCounselor(ctx echo.Context, request requests.CounselorRequest) (*domain.Counselors, []exceptions.ValidationMessage, error)
	AddSpecialist(ctx echo.Context, id uint, request requests.CounselorHasSpecialistRequest) ([]exceptions.ValidationMessage, error)
	DeleteSpecialist(ctx echo.Context, id uint, request requests.DeleteCounselorSpecialist) ([]exceptions.ValidationMessage, error)
}

type CounselorServiceImpl struct {
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

func (service *CounselorServiceImpl) AddSpecialist(ctx echo.Context, id uint, request requests.CounselorHasSpecialistRequest) ([]exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	counselor, errCounselor := service.CounselorRepo.FindById(id)
	if errCounselor != nil {
		return nil, errCounselor
	}

	specialist, errSpecialist := service.SpecialistRepo.FindSpecialistByName(request.Name)
	if errSpecialist != nil {
		return nil, errSpecialist
	}

	errAdd := service.CounselorHasSpecialistRepo.AddSpecialist(*counselor, specialist)
	if errAdd != nil {
		return nil, errAdd
	}

	return nil, nil
}

func (service *CounselorServiceImpl) DeleteSpecialist(ctx echo.Context, id uint, request requests.DeleteCounselorSpecialist) ([]exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	counselor, errCounselor := service.CounselorRepo.FindById(id)
	if errCounselor != nil {
		return nil, errCounselor
	}

	specialist, errSpecialist := service.SpecialistRepo.FindSpecialistByName(request.Name)
	if errSpecialist != nil {
		return nil, errSpecialist
	}

	preload, errPreload := service.CounselorRepo.PreloadSpecialist(counselor.Id)
	if errPreload != nil {
		return nil, errPreload
	}

	fmt.Println(preload)
	fmt.Println(specialist)

	// Verifikasi bahwa Specialist terkait dengan Counselor
	var isCategoryAssociated bool
	for _, c := range preload.Specialists {
		if c.Id == specialist.Id {
			fmt.Print(c.Id, specialist.Id)
			fmt.Println()
			isCategoryAssociated = true
			break
		}
	}

	if !isCategoryAssociated {
		return nil, errors.New("Specialist is not associated with the Counselor")
	}

	existingData := service.CounselorHasSpecialistRepo.DeleteSpecialistById(*counselor, specialist)
	if existingData != nil {
		return nil, existingData
	}

	return nil, nil
}
