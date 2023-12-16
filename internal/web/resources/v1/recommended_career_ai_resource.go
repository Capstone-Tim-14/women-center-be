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
