package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselorFavoriteRepository interface {
	AddCounselorFavorite(user domain.Users, counselor *domain.Counselors) error
	RemoveCounselorFavorite(user domain.Users, counselor []domain.Counselors) error
}

type CounselorFavoriteRepositoryImpl struct {
	Db *gorm.DB
}

func NewCounselorFavoriteRepository(db *gorm.DB) CounselorFavoriteRepository {
	return &CounselorFavoriteRepositoryImpl{
		Db: db,
	}
}

func (repository *CounselorFavoriteRepositoryImpl) AddCounselorFavorite(user domain.Users, counselor *domain.Counselors) error {
	result := repository.Db.Model(&user).Association("Counselors").Append(counselor)
	if result != nil {
		return result
	}

	return nil

}

func (repository *CounselorFavoriteRepositoryImpl) RemoveCounselorFavorite(user domain.Users, counselor []domain.Counselors) error {

	Transaction := repository.Db.Begin()

	result := Transaction.Model(&user).Association("Counselors").Delete(counselor)
	if result != nil {
		Transaction.Rollback()
		return result
	}

	Transaction.Commit()

	return nil
}
