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
	GetAllCounselors(ctx echo.Context) ([]domain.Counselors, error)
}

type CounselorServiceImpl struct {
	CounselorRepo repositories.CounselorRepository
	RoleRepo      repositories.RoleRepository
	validator     *validator.Validate
}

func NewCounselorService(counselor repositories.CounselorRepository, validator *validator.Validate, role repositories.RoleRepository) CounselorService {
	return &CounselorServiceImpl{
		CounselorRepo: counselor,
		RoleRepo:      role,
		validator:     validator,
	}
}

func (service *CounselorServiceImpl) RegisterCounselor(ctx echo.Context, request requests.CounselorRequest) (*domain.Counselors, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
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

func (service *CounselorServiceImpl) GetAllCounselors(ctx echo.Context) ([]domain.Counselors, error) {
	counselors, err := service.CounselorRepo.FindAllCounselors()
	if err != nil {
		return nil, fmt.Errorf("Error when get all counselors: %s", err.Error())
	}

	return counselors, nil
}
