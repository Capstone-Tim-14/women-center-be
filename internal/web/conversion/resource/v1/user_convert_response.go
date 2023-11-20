package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func UserDomainToUserResource(user *domain.Users) resources.UserResource {
	return resources.UserResource{
		Id:              user.Id,
		First_name:      user.First_name,
		Last_name:       user.Last_name,
		Email:           user.Credential.Email,
		Username:        user.Credential.Username,
		Profile_picture: user.Profile_picture,
		Phone_number:    user.Phone_number,
		Address:         user.Address,
		Status:          user.Status,
	}
}
