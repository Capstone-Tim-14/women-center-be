package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func UserDomainToAuthResource(user *domain.Users) resources.AuthResource {
	return resources.AuthResource{
		Id:       uint(user.Id),
		Fullname: user.First_name + user.Last_name,
		Username: user.Credential.Username,
		Role:     user.Credential.Role.Name,
		Email:    user.Email,
	}
}

func AuthResourceToAuthTokenResource(auth resources.AuthResource, token string) *resources.AuthTokenResource {
	return &resources.AuthTokenResource{
		Fullname: auth.Fullname,
		Email:    auth.Email,
		Token:    token,
	}
}
