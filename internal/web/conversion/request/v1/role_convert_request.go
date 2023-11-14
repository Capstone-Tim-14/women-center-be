package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func RoleCreateRequestToRoleDomain(request requests.RoleRequest) *domain.Roles {
	return &domain.Roles{
		Name: request.Name,
	}
}
