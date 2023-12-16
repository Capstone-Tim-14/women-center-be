package resources

type GenerateRecommendationCareerResource struct {
	Fullname string                                  `json:"fullname,omitempty"`
	Profile  string                                  `json:"profile,omitempty"`
	Chat     ChatHistoryRecommendationCareerResource `json:"chat,omitempty"`
}

type ChatHistoryRecommendationCareerResource struct {
	Question string `json:"question,omitempty"`
	Answer   string `json:"answer,omitempty"`
}

type HistoryRecommendationCareerAiResource struct {
	Id                uint                                      `json:"id,omitempty"`
	User_id           uint                                      `json:"user_id,omitempty"`
	Fullname          string                                    `json:"fullname,omitempty"`
	Profile           string                                    `json:"profile,omitempty"`
	HistoryChatCareer []ChatHistoryRecommendationCareerResource `json:"history_chat,omitempty"`
}
