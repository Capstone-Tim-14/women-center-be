package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func ConvertSaveHistoryChatToCareerRecommendationResource(user domain.Users, chat domain.HistoryRecommendationCareerAi) resources.GenerateRecommendationCareerResource {
	return resources.GenerateRecommendationCareerResource{
		Fullname: user.First_name + " " + user.Last_name,
		Profile:  user.Profile_picture,
		Chat: resources.ChatHistoryRecommendationCareerResource{
			Question: chat.User_question,
			Answer:   chat.Recommended_message,
		},
	}
}
