package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(counselor *domain.Counselors, scheduling []domain.Counseling_Schedule) error
	FindFreeSchedule(day, start, finish string) (*domain.Counseling_Schedule, error)
}

type ScheduleRepositoryImpl struct {
	Db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{
		Db: db,
	}
}

func (repository *ScheduleRepositoryImpl) CreateSchedule(counselor *domain.Counselors, scheduling []domain.Counseling_Schedule) error {

	transaction := repository.Db.Begin()

	err := transaction.Model(&counselor).Association("Schedules").Append(scheduling)

	if err != nil {
		transaction.Rollback()
		return err
	}

	transaction.Commit()

	return nil

}

func (repository *ScheduleRepositoryImpl) FindFreeSchedule(day, start, finish string) (*domain.Counseling_Schedule, error) {
	var Schedule domain.Counseling_Schedule

	DataSchedule := repository.Db.Preload("Counselors").Where("day = ? AND start = ? AND finish = ?", day, start, finish).First(&Schedule)
	if DataSchedule.Error != nil {
		return nil, DataSchedule.Error
	}

	return &Schedule, nil
}
