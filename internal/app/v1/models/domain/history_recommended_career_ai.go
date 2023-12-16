package domain

import (
	"time"

	"gorm.io/gorm"
)

type HistoryRecommendationCareerAi struct {
	Id                  uint `gorm:"primaryKey"`
	User_id             uint
	User_question       string
	Recommended_message string
	CreatedAt           time.Time      `gorm:"autoCreateTime"`
	UpdatedAt           time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

func (HistoryRecommendationCareerAi) TableName() string {
	return "history_user_recommended_career_ai"
}
