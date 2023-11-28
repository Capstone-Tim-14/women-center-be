package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func JobTypeCreateRequestToJobTypeDomain(request requests.JobTypeRequest) *domain.Job_Type {
	return &domain.Job_Type{
		Name:        request.Name,
		Description: request.Description,
	}
}
