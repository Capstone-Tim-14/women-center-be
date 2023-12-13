package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CounselingPackageRepository interface {
	CreatePackage(pack *domain.CounselingPackage) (*domain.CounselingPackage, error)
	FindByTitle(title string) ([]domain.CounselingPackage, error)
	FindById(id int) (*domain.CounselingPackage, error)
	GetAllPackage() ([]domain.CounselingPackage, error)
	DeletePackageById(id int) error
	UpdatePackage(id int, cpackage *domain.CounselingPackage) error
}

type CounselingPackageRepositoryImpl struct {
	db *gorm.DB
}

func NewCounselingPackageRepository(db *gorm.DB) CounselingPackageRepository {
	return &CounselingPackageRepositoryImpl{
		db: db,
	}
}

func (repository *CounselingPackageRepositoryImpl) FindById(id int) (*domain.CounselingPackage, error) {
	Package := domain.CounselingPackage{}

	result := repository.db.Where("id = ?", id).First(&Package)
	if result.Error != nil {
		return nil, result.Error
	}

	return &Package, nil
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

func (repository *CounselingPackageRepositoryImpl) DeletePackageById(id int) error {
	findPackage, err := repository.FindById(id)
	if err != nil {
		return fmt.Errorf("package not found")
	}

	result := repository.db.Delete(&findPackage)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *CounselingPackageRepositoryImpl) UpdatePackage(id int, cpackage *domain.CounselingPackage) error {
	result := repository.db.Model(&cpackage).Where("id = ?", id).Updates(cpackage)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
