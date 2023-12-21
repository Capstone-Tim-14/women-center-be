package handlers

import (
	"fmt"
	"strconv"
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselorHandler interface {
	RegisterHandler(ctx echo.Context) error
	AddSpecialist(ctx echo.Context) error
	RemoveManySpecialist(ctx echo.Context) error
	GetAllCounselorsHandler(echo.Context) error
	GetCounselorsForMobile(echo.Context) error
	GetDetailCounselorHandler(echo.Context) error
	GetDetailCounselorWeb(echo.Context) error
	GetCounselorProfile(ctx echo.Context) error
	UpdateCounselorHandler(echo.Context) error
	UpdateCounselorForMobile(echo.Context) error
	GetDetailCounselorSchedules(echo.Context) error
}

type CounselorHandlerImpl struct {
	CounselorService services.CounselorService
}

func NewCounselorHandler(counselor services.CounselorService) CounselorHandler {
	return &CounselorHandlerImpl{
		CounselorService: counselor,
	}
}

func (handler *CounselorHandlerImpl) RemoveManySpecialist(ctx echo.Context) error {
	var request requests.CounselorHasManyRequest

	GetId := ctx.Param("id")

	ErrBinding := ctx.Bind(&request)

	if ErrBinding != nil {
		return exceptions.StatusBadRequest(ctx, ErrBinding)
	}

	ParseToId, errParsing := strconv.Atoi(GetId)

	if errParsing != nil {
		fmt.Errorf(errParsing.Error())
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Invalid Format id"))
	}

	Validation, ErrGetSpecialist := handler.CounselorService.RemoveSpecialistCounselor(ctx, ParseToId, request)

	if Validation != nil || len(request.Name) == 0 {
		return exceptions.ValidationException(ctx, "Validation Error", Validation)
	}

	if ErrGetSpecialist != nil {
		if strings.Contains(ErrGetSpecialist.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, ErrGetSpecialist)
		}
		if strings.Contains(ErrGetSpecialist.Error(), "One of counselor request is not found") {
			return exceptions.StatusNotFound(ctx, ErrGetSpecialist)
		}

		return exceptions.StatusInternalServerError(ctx, ErrGetSpecialist)
	}

	return responses.StatusOK(ctx, "Counselor Remove Successfully", nil)
}

func (handler *CounselorHandlerImpl) RegisterHandler(ctx echo.Context) error {
	counselorCreateRequest := requests.CounselorRequest{}
	err := ctx.Bind(&counselorCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.CounselorService.RegisterCounselor(ctx, counselorCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exist") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorCreateResponse := conversion.CounselorDomainToCounselorResponse(response)

	return responses.StatusCreated(ctx, "Counselor created successfully", counselorCreateResponse)
}

func (handler *CounselorHandlerImpl) AddSpecialist(ctx echo.Context) error {
	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)
	var request requests.CounselorHasManyRequest
	errBinding := ctx.Bind(&request)
	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.CounselorService.AddSpecialist(ctx, uint(convertid), request)

	if validation != nil || len(request.Name) == 0 {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "One of Counselor request is not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success add specialist to counselor", nil)
}

func (handler *CounselorHandlerImpl) GetAllCounselorsHandler(ctx echo.Context) error {
	response, err := handler.CounselorService.GetAllCounselors(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorResponse := conversion.ConvertCounselorDomainToCounselorResponse(response)

	return responses.StatusOK(ctx, "Get all counselors successfully", counselorResponse)
}

func (handler *CounselorHandlerImpl) UpdateCounselorHandler(ctx echo.Context) error {
	counselorUpdateRequest := requests.UpdateCounselorProfileRequest{}
	picture, _ := ctx.FormFile("picture")
	err := ctx.Bind(&counselorUpdateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.CounselorService.UpdateCounselor(ctx, counselorUpdateRequest, picture)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Counselor updated successfully", nil)
}

func (handler *CounselorHandlerImpl) UpdateCounselorForMobile(ctx echo.Context) error {
	counselorUpdateRequest := requests.UpdateCounselorProfileRequestForMobile{}
	picture, _ := ctx.FormFile("picture")
	err := ctx.Bind(&counselorUpdateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.CounselorService.UpdateCounselorForMobile(ctx, counselorUpdateRequest, picture)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Counselor updated successfully", nil)
}

func (handler *CounselorHandlerImpl) GetCounselorsForMobile(ctx echo.Context) error {
	response, err := handler.CounselorService.GetCounselorsForMobile(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorResponse := conversion.ConvertCounselorDomainToCounselorResponse(response)

	return responses.StatusOK(ctx, "Get all counselors successfully", counselorResponse)
}

func (handler *CounselorHandlerImpl) GetDetailCounselorHandler(ctx echo.Context) error {

	response, err := handler.CounselorService.GetDetailCounselor(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Get detail counselor successfully", response)
}

func (handler *CounselorHandlerImpl) GetDetailCounselorSchedules(ctx echo.Context) error {

	response, err := handler.CounselorService.GetDetailCounselor(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Get detail counselor successfully", response)
}

func (handler *CounselorHandlerImpl) GetDetailCounselorWeb(ctx echo.Context) error {

	response, err := handler.CounselorService.GetDetailCounselorOnly(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Get detail counselor successfully", response)
}

func (handler *CounselorHandlerImpl) GetCounselorProfile(ctx echo.Context) error {
	response, err := handler.CounselorService.GetCounselorProfile(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	counselorResponse := conversion.CounselorDomainToProfileCounselor(response)

	return responses.StatusOK(ctx, "Get profile counselor successfully", counselorResponse)
}
