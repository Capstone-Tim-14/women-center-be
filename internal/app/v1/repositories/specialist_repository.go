package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type SpecialistRepository interface {
	CreateSpecialist(tag *domain.Specialist) (*domain.Specialist, error)
	// FindById(id int) (*domain.Specialist, error)
	FindSpecialistByName(name string) (*domain.Specialist, error)
	// FindAllTags() ([]domain.Specialist, error)
	// DeleteSpecialistById(tagID int) error
}

type SpecialistRepositoryImpl struct {
	db *gorm.DB
}

func NewSpecialistRepository(db *gorm.DB) SpecialistRepository {
	return &SpecialistRepositoryImpl{
		db: db,
	}
}

func (repository *SpecialistRepositoryImpl) CreateSpecialist(specialist *domain.Specialist) (*domain.Specialist, error) {
	result := repository.db.Create(&specialist)
	if result.Error != nil {
		return nil, result.Error
	}

	return specialist, nil
}

func (repository *SpecialistRepositoryImpl) FindSpecialistByName(name string) (*domain.Specialist, error) {
	specialist := domain.Specialist{}

	result := repository.db.Where("name = ?", name).First(&specialist)
	if result.Error != nil {
		return nil, result.Error
	}

	return &specialist, nil
}
