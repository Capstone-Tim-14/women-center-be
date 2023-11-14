package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func AdminDomainToAdminResource(admin *domain.Admin) resources.AdminResource {
	return resources.AdminResource{
		Id:         int(admin.Id),
		First_name: admin.First_name,
		Last_name:  admin.Last_name,
		Email:      admin.Email,
		Username:   admin.Credential.Username,
	}
}
