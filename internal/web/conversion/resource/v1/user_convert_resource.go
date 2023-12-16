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

func UserDomainToUsersResource(users []domain.Users) []resources.UserResource {

	var ResponseUser []resources.UserResource

	for _, item := range users {
		ResponseUser = append(ResponseUser, resources.UserResource{
			Id:         item.Id,
			First_name: item.First_name,
			Last_name:  item.Last_name,
			Email:      item.Credential.Email,
		})
	}

	return ResponseUser
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

func UserFavoriteArticleResponse(user *domain.Users) *[]resources.ArticleResource {
	userArticleFavorite := []resources.ArticleResource{}
	var authorName string
	var tagCategories []resources.ArticleCategory

	for _, artikelFav := range user.ArticleFavorites {

		if artikelFav.Admin != nil {
			authorName = artikelFav.Admin.First_name + " " + artikelFav.Admin.Last_name
		} else if artikelFav.Counselors != nil {
			authorName = artikelFav.Counselors.First_name + " " + artikelFav.Counselors.Last_name
		} else {
			authorName = "Unknown Author"
		}

		for _, tag := range artikelFav.Tags {
			tagCategories = append(tagCategories, resources.ArticleCategory{
				Id:   tag.Id,
				Name: tag.Name,
			})
		}

		favoriteArticle := resources.ArticleResource{
			Id:        artikelFav.Id,
			Title:     artikelFav.Title,
			Thumbnail: *artikelFav.Thumbnail,
			Slug:      artikelFav.Slug,
			Content:   artikelFav.Content,
			Status:    artikelFav.Status,
			Tag:       tagCategories,
			Author: resources.Author{
				Name: authorName,
			},
			PublishedAt: helpers.ParseOnlyDate(&artikelFav.PublishedAt),
		}
		userArticleFavorite = append(userArticleFavorite, favoriteArticle)
	}

	return &userArticleFavorite
}

func UserCounselorFavoriteResponse(user *domain.Users) []resources.CounselorResource {
	var Counselor []resources.CounselorResource
	for _, counselorFav := range user.Counselor_Favorite {
		Counselor = append(Counselor, resources.CounselorResource{
			Id:              counselorFav.Id,
			First_name:      counselorFav.First_name,
			Last_name:       counselorFav.Last_name,
			Email:           counselorFav.Credential.Email,
			Username:        counselorFav.Credential.Username,
			Profile_picture: counselorFav.Profile_picture,
			Description:     counselorFav.Description,
			Status:          counselorFav.Status,
		})
	}

	return Counselor
}
