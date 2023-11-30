package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func CounselorCreateRequestToCounselorDomain(request requests.CounselorRequest) *domain.Counselors {
	return &domain.Counselors{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Credential: &domain.Credentials{
			Username: request.Username,
			Password: request.Password,
			Role_id:  request.Role_id,
			Email:    request.Email,
		},
		Profile_picture: request.Profile_picture,
		Description:     request.Description,
		Status:          "ACTIVE",
	}
}

func CounselorUpdateRequestToCounselorDomain(request requests.CounselorRequest, counselor *domain.Counselors) *domain.Counselors {
	counselor.First_name = request.First_name
	counselor.Last_name = request.Last_name
	counselor.Credential.Username = request.Username
	counselor.Credential.Password = helpers.HashPassword(request.Password)
	counselor.Credential.Role_id = request.Role_id
	counselor.Credential.Email = request.Email
	counselor.Profile_picture = request.Profile_picture
	counselor.Description = request.Description

	return counselor
}
