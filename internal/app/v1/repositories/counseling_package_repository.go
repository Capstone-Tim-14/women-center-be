package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselingPackageRepository interface {
	CreatePackage(pack *domain.CounselingPackage) (*domain.CounselingPackage, error)
	FindByTitle(title string) ([]domain.CounselingPackage, error)
	GetAllPackage() ([]domain.CounselingPackage, error)
}

type CounselingPackageRepositoryImpl struct {
	db *gorm.DB
}

func NewCounselingRepository(db *gorm.DB) CounselingPackageRepository {
	return &CounselingPackageRepositoryImpl{
		db: db,
	}
}

func (repository *CounselingPackageRepositoryImpl) CreatePackage(pack *domain.CounselingPackage) (*domain.CounselingPackage, error) {
	result := repository.db.Create(&pack)
	if result.Error != nil {
		return nil, result.Error
	}

	return pack, nil

}

func (repository *CounselingPackageRepositoryImpl) FindByTitle(title string) ([]domain.CounselingPackage, error) {
	var counselingpackage []domain.CounselingPackage
	result := repository.db.Where("title = ?", title).First(&counselingpackage)
	if result.Error != nil {
		return nil, result.Error
	}

	return counselingpackage, nil
}

func (repository *CounselingPackageRepositoryImpl) GetAllPackage() ([]domain.CounselingPackage, error) {
	var lists []domain.CounselingPackage

	result := repository.db.Find(&lists)
	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Data Counseling Package is empty")
	}

	return lists, nil
}
