package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounseloHasSpecialistRepository interface {
	AddSpecialist(counselor domain.Counselors, specialist *domain.Specialist) error
	DeleteSpecialistById(counselor domain.Counselors, specialist *domain.Specialist) error
	RemoveManySpecialist(counselor domain.Counselors, specialist []domain.Specialist) error
}

type CounselorHasSpecialistRepositoryImpl struct {
	Db *gorm.DB
}

func NewCounselorHasSpecialistRepository(db *gorm.DB) CounseloHasSpecialistRepository {
	return &CounselorHasSpecialistRepositoryImpl{
		Db: db,
	}
}

func (repository *CounselorHasSpecialistRepositoryImpl) RemoveManySpecialist(counselor domain.Counselors, specialist []domain.Specialist) error {

	Transaction := repository.Db.Begin()

	result := Transaction.Model(&counselor).Association("Specialists").Delete(specialist)
	if result != nil {
		Transaction.Rollback()
		return result
	}

	Transaction.Commit()

	return nil
}

func (repository *CounselorHasSpecialistRepositoryImpl) AddSpecialist(counselor domain.Counselors, specialist *domain.Specialist) error {
	result := repository.Db.Model(&counselor).Association("Specialists").Append(specialist)
	if result != nil {
		return result
	}
	return nil
}

func (repository *CounselorHasSpecialistRepositoryImpl) DeleteSpecialistById(counselor domain.Counselors, specialist *domain.Specialist) error {
	result := repository.Db.Model(&counselor).Association("Specialists").Delete(specialist)
	if result != nil {
		return result
	}
	return nil
}
