package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func AdminCreateRequestToAdminDomain(request requests.AdminRequest) *domain.Admin {
	return &domain.Admin{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Credential: &domain.Credentials{
			Username: request.Username,
			Password: request.Password,
			Role_id:  request.Role_id,
			Email:    request.Email,
		},
	}
}
