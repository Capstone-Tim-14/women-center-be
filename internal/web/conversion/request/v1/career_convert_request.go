package conversion

import (
	"time"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func CareerCreateRequestToCareerDomain(request requests.CareerRequest) *domain.Career {
	return &domain.Career{
		Title_job:             request.Title_job,
		Logo:                  *request.Logo,
		Company_name:          request.Company_name,
		PublishedAt:           time.Now(),
		Cover:                 *request.Cover,
		Linkedin_url:          request.Linkedin_url,
		Location:              request.Location,
		About_job:             request.About_job,
		About_company:         request.About_company,
		Required_skill:        request.Required_skill,
		Company_industry:      request.Company_industry,
		Size_company_employee: request.Size_company_employee,
	}
}
