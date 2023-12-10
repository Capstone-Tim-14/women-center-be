package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func PackageCreateRequestToPackageDomain(request requests.CounselingPackageRequest) *domain.CounselingPackage {
	return &domain.CounselingPackage{
		Title:            request.Title,
		Thumbnail:        request.Thumbnail,
		Session_per_week: request.Session_per_week,
		Price:            helpers.StringToDecimal(request.Price),
		Description:      request.Description,
	}
}
