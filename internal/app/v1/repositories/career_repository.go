package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CareerRepository interface {
	CreateCareer(career *domain.Career) (*domain.Career, error)
	GetAllCareer() ([]domain.Career, error)
	FindCareerByid(id int) (*domain.Career, error)
	UpdateCareerById(id int, career *domain.Career) error
	DeleteCareerById(id int) error
	PreloadJobType(id uint) (*domain.Career, error)
}

type CareerRepositoryImpl struct {
	db *gorm.DB
}

func NewCareerRepository(db *gorm.DB) CareerRepository {
	return &CareerRepositoryImpl{
		db: db,
	}
}

func (repository *CareerRepositoryImpl) CreateCareer(career *domain.Career) (*domain.Career, error) {
	result := repository.db.Create(&career)
	if result.Error != nil {
		return nil, result.Error
	}

	return career, nil

}

func (repository *CareerRepositoryImpl) GetAllCareer() ([]domain.Career, error) {

	career := []domain.Career{}

	errTakeCareer := repository.db.Find(&career)

	if errTakeCareer.Error != nil {
		return nil, errTakeCareer.Error
	}

	if errTakeCareer.RowsAffected == 0 {
		return nil, errTakeCareer.Error
	}

	return career, nil

}

func (repository *CareerRepositoryImpl) FindCareerByid(id int) (*domain.Career, error) {

	career := domain.Career{}

	errTakeCareer := repository.db.Preload("Job_type").Where("id = ?", id).First(&career)

	if errTakeCareer.Error != nil {
		return nil, errTakeCareer.Error
	}

	if errTakeCareer.RowsAffected == 0 {
		return nil, errTakeCareer.Error
	}

	return &career, nil

}

func (repository *CareerRepositoryImpl) UpdateCareerById(id int, career *domain.Career) error {

	result := repository.db.Model(&career).Where("id = ?", id).Updates(&career)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (repository *CareerRepositoryImpl) DeleteCareerById(id int) error {

	career := domain.Career{}

	result := repository.db.Unscoped().Where("id = ?", id).Delete(&career)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (repository *CareerRepositoryImpl) PreloadJobType(id uint) (*domain.Career, error) {

	career := domain.Career{}
	if err := repository.db.Preload("Job_type").First(&career, id).Error; err != nil {
		return nil, err
	}

	return &career, nil
}
