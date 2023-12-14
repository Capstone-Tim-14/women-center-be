package handlers

import (
	"strconv"
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type CounselingPackageHandler interface {
	CreatePackage(ctx echo.Context) error
	FindByTitle(ctx echo.Context) error
	GetAllPackage(ctx echo.Context) error
	DeletePackage(ctx echo.Context) error
	UpdatePackage(ctx echo.Context) error
}

type CounselingPackageHandlerImpl struct {
	CounselingPackageService services.CounselingPackageService
}

func NewCounselingPackageHandler(cpackage services.CounselingPackageService) CounselingPackageHandler {
	return &CounselingPackageHandlerImpl{
		CounselingPackageService: cpackage,
	}
}

func (handler *CounselingPackageHandlerImpl) CreatePackage(ctx echo.Context) error {
	adminCreateRequest := requests.CounselingPackageRequest{}
	Thumbnail, errThumb := ctx.FormFile("thumbnail")

	if errThumb != nil {
		return exceptions.StatusBadRequest(ctx, errThumb)
	}

	err := ctx.Bind(&adminCreateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.CounselingPackageService.CreatePackage(ctx, adminCreateRequest, Thumbnail)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Title already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Package created succesfully", nil)
}

func (handler *CounselingPackageHandlerImpl) FindByTitle(ctx echo.Context) error {
	title := ctx.Param("title")
	response, err := handler.CounselingPackageService.FindByTitle(ctx, title)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	cpackageResponse := conversion.ConvertCounselingPackageDomainToResponse(response)

	return responses.StatusOK(ctx, "Succes Get Package", cpackageResponse)
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

func (handler *CounselingPackageHandlerImpl) DeletePackage(ctx echo.Context) error {
	pkgId := ctx.Param("id")
	pkgIdInt, errId := strconv.Atoi(pkgId)
	if errId != nil {
		return exceptions.StatusBadRequest(ctx, errId)
	}

	errPkg := handler.CounselingPackageService.DeletePackageById(ctx, pkgIdInt)
	if errPkg != nil {
		if strings.Contains(errPkg.Error(), "package not found") {
			return exceptions.StatusNotFound(ctx, errPkg)
		}
		return exceptions.StatusInternalServerError(ctx, errPkg)
	}

	return responses.StatusOK(ctx, "Success remove package", nil)
}

func (handler *CounselingPackageHandlerImpl) UpdatePackage(ctx echo.Context) error {
	var request requests.CounselingPackageRequest
	errBinding := ctx.Bind(&request)

	Thumbnail, _ := ctx.FormFile("thumbnail")

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	_, validation, err := handler.CounselingPackageService.UpdatePackageById(ctx, request, Thumbnail)
	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "invalid id") {
			return exceptions.StatusBadRequest(ctx, err)
		}
		if strings.Contains(err.Error(), "package not found") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Package updated!", nil)
}
