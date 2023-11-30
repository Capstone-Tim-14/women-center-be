package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselorRepository interface {
	CreateCounselor(counselor *domain.Counselors) (*domain.Counselors, error)
	UpdateCounselor(counselor *domain.Counselors, id int) (*domain.Counselors, error)
	FindyByEmail(email string) (*domain.Counselors, error)
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

func (repository *CounselorRepositoryImpl) FindAllCounselors() ([]domain.Counselors, error) {
	counselor := []domain.Counselors{}

	result := repository.db.Preload("Credential").Preload("Credential.Role").Find(&counselor)
	if result.Error != nil {
		return nil, result.Error
	}

	return counselor, nil
}

func (repository *CounselorRepositoryImpl) UpdateCounselor(counselor *domain.Counselors, id int) (*domain.Counselors, error) {
	result := repository.db.Model(&counselor).Where("id = ?", id).Updates(counselor)
	if result.Error != nil {
		return nil, result.Error
	}

	return counselor, nil
}
