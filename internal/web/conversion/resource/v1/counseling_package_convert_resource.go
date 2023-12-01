package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func ConvertCounselingPackageDomainToResponse(cpackage []domain.CounselingPackage) []resources.CounselingPackageResource {
	listResponse := []resources.CounselingPackageResource{}

	for _, val := range cpackage {
		listResponse = append(listResponse, resources.CounselingPackageResource{
			Id:                 val.Id,
			Thumbnail:          *val.Thumbnail,
			Package_name:       val.Package_name,
			Description:        val.Description,
			Number_of_sessions: val.Number_of_sessions,
			Price:              val.Price.String(),
			PublishedAt:        helpers.ParseDateFormat(&val.PublishedAt),
		})
	}

	return listResponse
}
