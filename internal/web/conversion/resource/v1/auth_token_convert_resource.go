package conversion

import "woman-center-be/internal/web/resources/v1"

func AuthResourceToAuthTokenResource(auth resources.AuthResource, token string) *resources.AuthTokenResource {
	return &resources.AuthTokenResource{
		Fullname: auth.Fullname,
		Email:    auth.Email,
		Token:    token,
	}
}
