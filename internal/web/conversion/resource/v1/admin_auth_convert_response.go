package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func AdminDomainToAuthResource(user *domain.Admin) resources.AuthResource {
	return resources.AuthResource{
		Id:       uint(user.Id),
		Fullname: user.First_name + " " + user.Last_name,
		Username: user.Credential.Username,
		Role:     user.Credential.Role.Name,
		Email:    user.Credential.Email,
	}
}
