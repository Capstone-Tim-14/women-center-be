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
	UserProfile := resources.GetUserProfile{
		Id:              user.Id,
		Profile_picture: user.Profile_picture,
		Username:        user.Credential.Username,
		Full_name:       user.First_name + " " + user.Last_name,
		Email:           user.Credential.Email,
	}

	if user.Birthday != nil {
		UserProfile.Birthday = helpers.ParseOnlyDate(user.Birthday)
	}

	return UserProfile
}

func UserDomainToUserUpdateProfileResource(user *domain.Users) resources.UpdateUserProfile {
	return resources.UpdateUserProfile{
		Id:              user.Id,
		First_name:      user.First_name,
		Last_name:       user.Last_name,
		Username:        user.Credential.Username,
		Email:           user.Credential.Email,
		Birthday:        helpers.ParseOnlyDate(user.Birthday),
		Profile_picture: user.Profile_picture,
	}
}

func UserDomainToUserOTPGenerate(code string, secret string) *resources.OtpResources {
	return &resources.OtpResources{
		Code:   code,
		Secret: secret,
	}
}

func UserCounselorFavoriteResponse(user *domain.Users) resources.UserCounselorFavorite {
	var Counselor []resources.CounselorFavorite
	CounselorFavoriteResource := resources.UserCounselorFavorite{}
	CounselorFavoriteResource.Id = user.Id
	CounselorFavoriteResource.First_name = user.First_name
	CounselorFavoriteResource.Last_name = user.Last_name
	CounselorFavoriteResource.Username = user.Credential.Username
	for _, counselorFav := range user.Counselor_Favorite {
		Counselor = append(Counselor, resources.CounselorFavorite{
			Id:              counselorFav.Id,
			First_name:      counselorFav.First_name,
			Last_name:       counselorFav.Last_name,
			Username:        counselorFav.Credential.Username,
			Profile_picture: counselorFav.Profile_picture,
		})
	}
	CounselorFavoriteResource.CounselorFavorite = Counselor

	return CounselorFavoriteResource
}
