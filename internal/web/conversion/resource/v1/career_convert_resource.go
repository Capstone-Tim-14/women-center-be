package conversion

import (
	"strconv"
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
		singleCareerResource.Recomendation = strconv.FormatBool(career.Recomendation)
		singleCareerResource.PublishedAt = helpers.ParseDateFormat(&career.PublishedAt)
		careerResource = append(careerResource, singleCareerResource)
	}

	return careerResource
}

func ConvertCareerDetailResource(career *domain.Career) *resources.CareerResource {
	DetailJobType := []resources.JobType{}
	careerDetailResource := resources.CareerResource{}
	careerDetailResource.Title_job = career.Title_job
	careerDetailResource.Company_name = career.Company_name
	careerDetailResource.Logo = career.Logo
	careerDetailResource.Location = career.Location
	careerDetailResource.PublishedAt = helpers.ParseDateFormat(&career.PublishedAt)
	careerDetailResource.Cover = career.Cover
	careerDetailResource.Linkedin_url = career.Linkedin_url
	careerDetailResource.About_job = career.About_job
	careerDetailResource.About_company = career.About_company
	careerDetailResource.Required_skill = career.Required_skill
	careerDetailResource.Company_industry = career.Company_industry
	careerDetailResource.Size_company_employee = career.Size_company_employee
	for _, jobType := range career.Job_type {
		DetailJobType = append(DetailJobType, resources.JobType{
			Name:        *&jobType.Name,
			Description: jobType.Description,
		})
	}
	careerDetailResource.Job_type = DetailJobType
	careerDetailResource.CreatedAt = helpers.ParseDateFormat(&career.CreatedAt)
	careerDetailResource.UpdatedAt = helpers.ParseDateFormat(&career.UpdatedAt)

	return &careerDetailResource
}

func ConvertRecomendationCareer(careers []domain.Career) []resources.CareerResource {
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
