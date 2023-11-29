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
		singleCareerResource.Id = career.Id
		singleCareerResource.Title_job = career.Title_job
		singleCareerResource.Company_name = career.Company_name
		singleCareerResource.Logo = career.Logo
		singleCareerResource.Location = career.Location
		singleCareerResource.PublishedAt = helpers.ParseDateFormat(&career.PublishedAt)
		careerResource = append(careerResource, singleCareerResource)
	}

	return careerResource
}

func ConvertCareerDetailResource(career *domain.Career) *resources.CareerResource {
	careerDetailResource := &resources.CareerResource{
		Title_job:             career.Title_job,
		Company_name:          career.Company_name,
		Logo:                  career.Logo,
		Cover:                 career.Cover,
		Required_skill:        career.Required_skill,
		Size_company_employee: career.Size_company_employee,
		About_job:             career.About_job,
		About_company:         career.About_company,
		PublishedAt:           helpers.ParseDateFormat(&career.PublishedAt),
	}

	return careerDetailResource
}
