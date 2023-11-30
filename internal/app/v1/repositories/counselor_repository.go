package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselorRepository interface {
	CreateCounselor(counselor *domain.Counselors) (*domain.Counselors, error)
	FindyByEmail(email string) (*domain.Counselors, error)
	FindById(id uint) (*domain.Counselors, error)
	FindAllCounselors() ([]domain.Counselors, error)
}

type CounselorRepositoryImpl struct {
	db *gorm.DB
}

func NewCounselorRepository(db *gorm.DB) CounselorRepository {
	return &CounselorRepositoryImpl{
		db: db,
	}
}

func (repository *CounselorRepositoryImpl) CreateCounselor(counselor *domain.Counselors) (*domain.Counselors, error) {
	result := repository.db.Create(&counselor)
	if result.Error != nil {
		return nil, result.Error
	}

	return counselor, nil
}

func (repository *CounselorRepositoryImpl) FindyByEmail(email string) (*domain.Counselors, error) {
	counselor := domain.Counselors{}

	result := repository.db.InnerJoins("Credential").InnerJoins("Credential.Role").Where("email = ?", email).First(&counselor)
	if result.Error != nil {
		return nil, result.Error
	}

	return &counselor, nil
}

func (repository *CounselorRepositoryImpl) FindById(id uint) (*domain.Counselors, error) {
	counselor := domain.Counselors{}

	result := repository.db.Preload("Specialists").Where("id = ?", id).First(&counselor)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Counselor not found")
	}
	return &counselor, nil
}
func (repository *CounselorRepositoryImpl) FindAllCounselors() ([]domain.Counselors, error) {
	counselor := []domain.Counselors{}

	result := repository.db.Preload("Credential").Preload("Credential.Role").Find(&counselor)
	if result.Error != nil {
		return nil, result.Error
	}

	return counselor, nil
}
