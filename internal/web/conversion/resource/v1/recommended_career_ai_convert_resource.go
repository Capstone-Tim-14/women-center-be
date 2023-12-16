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

func ConvertHistoryChatToCareerRecommendationResource(user *domain.Users) resources.HistoryRecommendationCareerAiResource {

	historyChat := []resources.ChatHistoryRecommendationCareerResource{}
	chatResource := resources.HistoryRecommendationCareerAiResource{}
	chatResource.Fullname = user.First_name + " " + user.Last_name
	chatResource.Profile = user.Profile_picture
	for _, chat := range user.HistoryChatRecommendationCareer {
		historyChat = append(historyChat, resources.ChatHistoryRecommendationCareerResource{
			Question: chat.User_question,
			Answer:   chat.Recommended_message,
		})
	}
	chatResource.HistoryChatCareer = historyChat

	return chatResource
}
