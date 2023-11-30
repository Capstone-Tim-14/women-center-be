package services

import (
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
	AddSpecialist(ctx echo.Context, id uint, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error)
	RemoveSpecialistCounselor(ctx echo.Context, id int, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error)
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

func (service *CounselorServiceImpl) RemoveSpecialistCounselor(ctx echo.Context, id int, request requests.CounselorHasManyRequest) ([]exceptions.ValidationMessage, error) {
	ValidationMessage := service.Validator.Struct(request)
	var Specialist []domain.Specialist

	if ValidationMessage != nil || len(request.Name) == 0 {
		return helpers.ValidationError(ctx, ValidationMessage), nil
	}

	GetCounselorData, errGetCounselor := service.CounselorRepo.FindById(uint(id))

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

	counselor, errCounselor := service.CounselorRepo.FindById(id)
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
