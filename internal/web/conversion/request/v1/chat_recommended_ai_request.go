package conversion

import "woman-center-be/internal/app/v1/models/domain"

func ConvertChatToHistoryRecommendationCareer(user domain.Users, question string, message string) domain.HistoryRecommendationCareerAi {
	return domain.HistoryRecommendationCareerAi{
		User_id:             user.Id,
		User_question:       question,
		Recommended_message: message,
	}
}
