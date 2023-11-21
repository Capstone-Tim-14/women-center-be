package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func RoleDomainToRoleResource(role *domain.Roles) resources.RoleResource {
	return resources.RoleResource{
		Id:   role.Id,
		Name: role.Name,
	}
}

func ConvertRoleResource(roles []domain.Roles) []resources.RoleResource {
	roleResource := []resources.RoleResource{}
	for _, role := range roles {
		roleResource = append(roleResource, resources.RoleResource{
			Id:   role.Id,
			Name: role.Name,
		})
	}

	return roleResource
}
