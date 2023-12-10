package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselorFavoriteRepository interface {
	AddCounselorFavorite(user domain.Users, counselor *domain.Counselors) error
	RemoveCounselorFavorite(user domain.Users, counselor *domain.Counselors) error
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
	result := repository.Db.Model(&user).Association("Counselor_Favorite").Append(counselor)
	if result != nil {
		return result
	}

	return nil

}

func (repository *CounselorFavoriteRepositoryImpl) RemoveCounselorFavorite(user domain.Users, counselor *domain.Counselors) error {
	transaction := repository.Db.Begin()

	result := transaction.Model(&user).Association("Counselor_Favorite").Delete(counselor)
	if result != nil {
		transaction.Rollback()
		return result
	}

	transaction.Commit()

	return nil
}
