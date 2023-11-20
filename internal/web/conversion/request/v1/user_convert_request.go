package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func UserCreateRequestToUserDomain(request requests.UserRequest) *domain.Users {
	return &domain.Users{
		First_name: request.First_name,
		Last_name:  request.Last_name,
		Credential: &domain.Credentials{
			Email:    request.Email,
			Username: request.Username,
			Password: request.Password,
			Role_id:  request.Role_id,
		},
		Profile_picture: request.Profile_picture,
		Phone_number:    request.Phone_number,
		Address:         request.Address,
		Status:          "ACTIVE",
	}
}
