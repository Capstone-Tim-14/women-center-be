package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func JobTypeDomainToJobTypeResponse(jobtype *domain.Job_Type) resources.JobTypeResource {
	return resources.JobTypeResource{
		Id:          jobtype.Id,
		Name:        jobtype.Name,
		Description: jobtype.Description,
	}
}

func ConvertJobTypeResource(jobtypes []domain.Job_Type) []resources.JobTypeResource {
	jobtypeResource := []resources.JobTypeResource{}
	for _, jobtype := range jobtypes {
		jobtypeResource = append(jobtypeResource, resources.JobTypeResource{
			Id:          jobtype.Id,
			Name:        jobtype.Name,
			Description: jobtype.Description,
		})
	}

	return jobtypeResource
}
