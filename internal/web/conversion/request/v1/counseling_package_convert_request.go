package conversion

import (
	"time"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func PackageCreateRequestToPackageDomain(request requests.CounselingPackageRequest) *domain.CounselingPackage {
	return &domain.CounselingPackage{
		Package_name:       request.Package_name,
		Description:        request.Description,
		PublishedAt:        time.Now(),
		Thumbnail:          request.Thumbnail,
		Number_of_sessions: request.Number_of_sessions,
		Price:              helpers.StringToDecimal(request.Price),
	}
}
