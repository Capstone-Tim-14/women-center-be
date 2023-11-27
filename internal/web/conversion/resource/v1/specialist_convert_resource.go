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

func ConvertSpecialistResource(lists []domain.Specialist) []resources.SpecialistResource {
	listResource := []resources.SpecialistResource{}
	for _, list := range lists {
		listResource = append(listResource, resources.SpecialistResource{
			Id:          list.Id,
			Name:        list.Name,
			Description: list.Description,
		})
	}

	return listResource
}
