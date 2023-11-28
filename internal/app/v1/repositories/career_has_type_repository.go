package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CareerhasTypeRepository interface {
	AddJobType(career domain.Career, jobtype *domain.Job_Type) error
	RemoveJobTypeById(career domain.Career, jobtype *domain.Job_Type) error
}

type CareerhasTypeRepositoryImpl struct {
	Db *gorm.DB
}

func NewCareerhasTypeRepository(db *gorm.DB) CareerhasTypeRepository {
	return &CareerhasTypeRepositoryImpl{
		Db: db,
	}
}

func (repository *CareerhasTypeRepositoryImpl) AddJobType(career domain.Career, jobtype *domain.Job_Type) error {
	result := repository.Db.Model(&career).Association("Job_type").Append(jobtype)
	if result != nil {
		return result
	}
	return nil
}

func (repository *CareerhasTypeRepositoryImpl) RemoveJobTypeById(career domain.Career, jobtype *domain.Job_Type) error {
	result := repository.Db.Model(&career).Association("Job_type").Delete(jobtype)
	if result != nil {
		return result
	}
	return nil
}
