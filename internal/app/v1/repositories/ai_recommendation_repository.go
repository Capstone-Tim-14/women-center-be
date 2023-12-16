package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type RecommendationAiRepository interface {
	SaveHistoryRecommendationCareer(domain.HistoryRecommendationCareerAi) (*domain.HistoryRecommendationCareerAi, error)
	FindAllHistoryRecommendationCareer() (*domain.Users, error)
}

type NewRecommendationAiRepositoryImpl struct {
	Db *gorm.DB
}

func NewRecommendationAiRepository(db *gorm.DB) RecommendationAiRepository {
	return &NewRecommendationAiRepositoryImpl{
		Db: db,
	}
}

func (repo *NewRecommendationAiRepositoryImpl) SaveHistoryRecommendationCareer(chat domain.HistoryRecommendationCareerAi) (*domain.HistoryRecommendationCareerAi, error) {

	errCreateData := repo.Db.Create(&chat)

	if errCreateData.Error != nil {
		return nil, fmt.Errorf("Error save chat")
	}

	return &chat, nil

}

func (repo *NewRecommendationAiRepositoryImpl) FindAllHistoryRecommendationCareer() (*domain.Users, error) {
	var user domain.Users

	err := repo.Db.Preload("history_user_recommended_career_ai").Find(&user).Error

	if err.Error != nil {
		return nil, fmt.Errorf("Error get all history chat")
	}

	return &user, nil
}
