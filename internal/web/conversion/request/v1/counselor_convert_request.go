package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
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
		Phone_number:    request.Phone_number,
		Description:     request.Description,
		Status:          "ACTIVE",
	}
}
