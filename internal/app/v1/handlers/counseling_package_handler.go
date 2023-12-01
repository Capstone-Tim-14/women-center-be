package handlers

import "woman-center-be/internal/app/v1/services"

type CounselingPackageHandler interface {
}

type CounselingPackageHandlerImpl struct {
	CounselingPackageService services.CounselingPackageService
}

func NewCounselingPackageHandler(cpackage services.CounselingPackageService) CounselingPackageHandler {
	return &CounselingPackageHandlerImpl{
		CounselingPackageService: cpackage,
	}
}
