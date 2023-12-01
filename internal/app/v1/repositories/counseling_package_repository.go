package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselingPackageRepository interface {
	CreatePackage(pack *domain.CounselingPackage) (*domain.CounselingPackage, error)
	FindByName(name string) (*domain.CounselingPackage, error)
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

func (repository *CounselingPackageRepositoryImpl) FindByName(name string) (*domain.CounselingPackage, error) {
	counselingpackage := domain.CounselingPackage{}
	result := repository.db.Where("package_name = ?", name).First(&counselingpackage)
	if result.Error != nil {
		return nil, result.Error
	}

	return &counselingpackage, nil
}
