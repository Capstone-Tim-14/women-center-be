package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func SpecialistCreateRequestToSpecialistDomain(request requests.SpecialistRequest) *domain.Specialist {
	return &domain.Specialist{
		Name:        request.Name,
		Description: request.Description,
	}
}
