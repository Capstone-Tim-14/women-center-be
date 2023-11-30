package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"

	"gorm.io/gorm"
)

type CareerRepository interface {
	CreateCareer(career *domain.Career) (*domain.Career, error)
	GetAllCareer(job requests.CareerFilterRequest) ([]domain.Career, error)
	FindCareerByid(id int) (*domain.Career, error)
	UpdateCareerById(id int, career *domain.Career) error
	DeleteCareerById(id int) error
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

func (repository *CareerRepositoryImpl) GetAllCareer(job requests.CareerFilterRequest) ([]domain.Career, error) {

	var career []domain.Career

	var errTakeCareer *gorm.DB

	if len(job.JobType) > 0 {
		errTakeCareer = repository.db.Joins("INNER JOIN career_has_types ON careers.id = career_has_types.career_id").Joins("INNER JOIN job_types ON career_has_types.job_type_id = job_types.id").Where("job_types.name IN (?)", job.JobType).Distinct().Find(&career)
	} else {
		errTakeCareer = repository.db.Find(&career)
	}

	if errTakeCareer.Error != nil {
		return nil, errTakeCareer.Error
	}

	if errTakeCareer.RowsAffected == 0 {
		return nil, fmt.Errorf("Career is empty")
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
