package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CareerRepository interface {
	CreateCareer(career *domain.Career) (*domain.Career, error)
}

type CareerRepositoryImpl struct {
	db *gorm.DB
}

func NewCareerRepository(db *gorm.DB) CareerRepository {
	return &CareerRepositoryImpl{
		db: db,
	}
}

func (repository *CareerRepositoryImpl) CreateCareer(career *domain.Career) (*domain.Career, error) {
	result := repository.db.Create(&career)
	if result.Error != nil {
		return nil, result.Error
	}

	return career, nil

}
