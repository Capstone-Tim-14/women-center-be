package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CareerRepository interface {
	CreateCareer(career *domain.Career) (*domain.Career, error)
	GetAllCareer() ([]domain.Career, error)
	FindCareerByid(id int) (*domain.Career, error)
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

func (repository *CareerRepositoryImpl) GetAllCareer() ([]domain.Career, error) {

	career := []domain.Career{}

	errTakeCareer := repository.db.Find(&career)

	if errTakeCareer.Error != nil {
		return nil, errTakeCareer.Error
	}

	if errTakeCareer.RowsAffected == 0 {
		return nil, errTakeCareer.Error
	}

	return career, nil

}

func (repository *CareerRepositoryImpl) FindCareerByid(id int) (*domain.Career, error) {

	career := domain.Career{}

	errTakeCareer := repository.db.Where("id = ?", id).First(&career)

	if errTakeCareer.Error != nil {
		return nil, errTakeCareer.Error
	}

	if errTakeCareer.RowsAffected == 0 {
		return nil, errTakeCareer.Error
	}

	return &career, nil

}
