package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func CounselorDomainToCounselorResponse(counselor *domain.Counselors) resources.CounselorResource {
	return resources.CounselorResource{
		Id:              counselor.Id,
		First_name:      counselor.First_name,
		Last_name:       counselor.Last_name,
		Email:           counselor.Email,
		Username:        counselor.Credential.Username,
		Profile_picture: counselor.Profile_picture,
		Phone_number:    counselor.Phone_number,
		Address:         counselor.Address,
		Status:          counselor.Status,
	}
}

func RoleDomainToRoleResponse(role *domain.Roles) resources.RoleResource {
	return resources.RoleResource{
		Id:   role.Id,
		Name: role.Name,
	}
}
