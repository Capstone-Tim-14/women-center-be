package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type JobTypeRepository interface {
	CreateJobType(tag *domain.Job_Type) (*domain.Job_Type, error)
	FindJobTypeByName(name string) (*domain.Job_Type, error)
	ShowAllJobType() ([]domain.Job_Type, error)
}

type JobTypeRepositoryImpl struct {
	db *gorm.DB
}

func NewJobTypeRepository(db *gorm.DB) JobTypeRepository {
	return &JobTypeRepositoryImpl{
		db: db,
	}
}

func (repository *JobTypeRepositoryImpl) CreateJobType(jobtype *domain.Job_Type) (*domain.Job_Type, error) {
	result := repository.db.Create(&jobtype)
	if result.Error != nil {
		return nil, result.Error
	}

	return jobtype, nil
}

func (repository *JobTypeRepositoryImpl) FindJobTypeByName(name string) (*domain.Job_Type, error) {
	jobtype := domain.Job_Type{}

	result := repository.db.Where("name = ?", name).First(&jobtype)
	if result.Error != nil {
		return nil, result.Error
	}

	return &jobtype, nil
}

func (repository *JobTypeRepositoryImpl) ShowAllJobType() ([]domain.Job_Type, error) {
	var types []domain.Job_Type

	result := repository.db.Find(&types)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("Job types is empty")
	}

	return types, nil
}
