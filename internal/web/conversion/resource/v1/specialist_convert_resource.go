package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func SpecialistDomainToSpecialistResponse(specialist *domain.Specialist) resources.SpecialistResource {
	return resources.SpecialistResource{
		Id:          specialist.Id,
		Name:        specialist.Name,
		Description: specialist.Description,
	}
}
