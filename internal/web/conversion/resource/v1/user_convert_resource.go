package conversion

import (
	"time"
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
	userBirthday := user.Birthday
	convertBirthday := TimeToStringFormat(userBirthday)

	return resources.UpdateUserProfile{
		Id:              user.Id,
		First_name:      user.First_name,
		Last_name:       user.Last_name,
		Username:        user.Credential.Username,
		Email:           user.Credential.Email,
		Birthday:        convertBirthday,
		Profile_picture: user.Profile_picture,
	}
}

func TimeToStringFormat(t *time.Time) string {
	return t.Format("2006-01-02") // Mengubah time.Time menjadi string dengan format "YYYY-MM-DD"
}
