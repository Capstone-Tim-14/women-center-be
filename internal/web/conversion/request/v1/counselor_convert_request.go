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

func CounselorUpdateRequestToCounselorDomain(request requests.CounselorRequest) *domain.Counselors {
	return &domain.Counselors{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Credential: &domain.Credentials{
			Username: request.Username,
			Password: helpers.HashPassword(request.Password),
			Email:    request.Email,
			Role_id:  request.Role_id,
		},
		Profile_picture: request.Profile_picture,
		Description:     request.Description,
	}
}
