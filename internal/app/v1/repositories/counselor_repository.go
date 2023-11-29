package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselorRepository interface {
	CreateCounselor(counselor *domain.Counselors) (*domain.Counselors, error)
	FindById(id int) (*domain.Counselors, error)
	FindyByEmail(email string) (*domain.Counselors, error)
}

type CounselorRepositoryImpl struct {
	db *gorm.DB
}

func NewCounselorRepository(db *gorm.DB) CounselorRepository {
	return &CounselorRepositoryImpl{
		db: db,
	}
}

func (repository *CounselorRepositoryImpl) FindById(id int) (*domain.Counselors, error) {

	var Counselor domain.Counselors

	ErrGetCounselor := repository.db.First(&Counselor, "id = ?", id)

	if ErrGetCounselor.Error != nil {
		return nil, fmt.Errorf("Counselor not found")
	}

	return &Counselor, nil

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
