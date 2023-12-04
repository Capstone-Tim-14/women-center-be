package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func ConvertCounselingPackageDomainToResponse(cpackage []domain.CounselingPackage) []resources.CounselingPackageResource {
	listResponse := []resources.CounselingPackageResource{}

	for _, val := range cpackage {
		listResponse = append(listResponse, resources.CounselingPackageResource{
			Id:               val.Id,
			Thumbnail:        *val.Thumbnail,
			Title:            val.Title,
			Session_per_week: val.Session_per_week,
			Price:            val.Price.String(),
			Description:      val.Description,
		})
	}

	return listResponse
}
