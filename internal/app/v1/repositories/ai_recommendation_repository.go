package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type RecommendationAiRepository interface {
	SaveHistoryRecommendationCareer(domain.HistoryRecommendationCareerAi) (*domain.HistoryRecommendationCareerAi, error)
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
