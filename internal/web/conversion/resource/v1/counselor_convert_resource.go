package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func CounselorDomainToCounselorResponse(counselor *domain.Counselors) resources.CounselorResource {
	return resources.CounselorResource{
		Id:           counselor.Id,
		First_name:   counselor.First_name,
		Last_name:    counselor.Last_name,
		Email:        counselor.Credential.Email,
		Phone_number: counselor.Phone_number,
		Description:  counselor.Description,
		Status:       counselor.Status,
	}
}
