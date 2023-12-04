package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselingPackageHandler interface {
	GetAllPackage(ctx echo.Context) error
}

type CounselingPackageHandlerImpl struct {
	CounselingPackageService services.CounselingPackageService
}

func NewCounselingPackageHandler(cpackage services.CounselingPackageService) CounselingPackageHandler {
	return &CounselingPackageHandlerImpl{
		CounselingPackageService: cpackage,
	}
}

func (handler *CounselingPackageHandlerImpl) GetAllPackage(ctx echo.Context) error {
	response, err := handler.CounselingPackageService.GetAllPackage(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Data Counseling Package is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	listResponse := conversion.ConvertCounselingPackageDomainToResponse(response)

	return responses.StatusOK(ctx, "Success Get All Counseling Package", listResponse)
}
