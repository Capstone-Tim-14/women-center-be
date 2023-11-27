package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func ConvertCareerRsource(careers []domain.Career) []resources.CareerResource {
	careerResource := []resources.CareerResource{}
	for _, career := range careers {
		singleCareerResource := resources.CareerResource{}
		singleCareerResource.Title_job = career.Title_job
		singleCareerResource.Company_name = *&career.Company_name
		singleCareerResource.Logo = career.Logo
		singleCareerResource.Location = *&career.Location
		singleCareerResource.PublishedAt = helpers.ParseDateFormat(&career.PublishedAt)
		careerResource = append(careerResource, singleCareerResource)
	}

	return careerResource
}
