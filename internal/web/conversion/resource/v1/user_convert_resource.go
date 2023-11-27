package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
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
		Status:          user.Status,
	}
}

func UserDomainToUserProfileResource(user *domain.Users) resources.GetUserProfile {
	return resources.GetUserProfile{
		Id:              user.Id,
		Profile_picture: user.Profile_picture,
		Username:        user.Credential.Username,
		Full_name:       user.First_name + " " + user.Last_name,
		Email:           user.Credential.Email,
	}
}

func UserDomainToUserUpdateProfileResource(user *domain.Users) resources.UpdateUserProfile {
	return resources.UpdateUserProfile{
		Id:              user.Id,
		First_name:      user.First_name,
		Last_name:       user.Last_name,
		Username:        user.Credential.Username,
		Email:           user.Credential.Email,
		Birthday:        helpers.ParseDateFormat(user.Birthday),
		Profile_picture: user.Profile_picture,
	}
}
