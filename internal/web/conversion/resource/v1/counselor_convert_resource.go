package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func CounselorDomainToCounselorResponse(counselor *domain.Counselors) resources.CounselorResource {
	return resources.CounselorResource{
		Id:              counselor.Id,
		First_name:      counselor.First_name,
		Last_name:       counselor.Last_name,
		Email:           counselor.Credential.Email,
		Description:     counselor.Description,
		Status:          counselor.Status,
		Profile_picture: counselor.Profile_picture,
	}
}

func ConvertCounselorDomainToCounselorResponse(counselor []domain.Counselors) []resources.CounselorResource {
	counselorResponse := []resources.CounselorResource{}

	for _, counselor := range counselor {
		counselorResponse = append(counselorResponse, resources.CounselorResource{
			Id:              counselor.Id,
			First_name:      counselor.First_name,
			Last_name:       counselor.Last_name,
			Email:           counselor.Credential.Email,
			Description:     counselor.Description,
			Status:          counselor.Status,
			Profile_picture: counselor.Profile_picture,
		})
	}

	return counselorResponse
}
