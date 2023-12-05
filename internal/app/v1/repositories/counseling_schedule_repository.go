package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(counselor *domain.Counselors, scheduling []domain.Counseling_Schedule) error
	CheckDayCounselingScheduleExists(id int, day string) (*domain.Counseling_Schedule, error)
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

func (repository *ScheduleRepositoryImpl) CheckDayCounselingScheduleExists(id int, day string) (*domain.Counseling_Schedule, error) {

	var counselorSchedule domain.Counseling_Schedule

	errGetDaySchedule := repository.Db.Where("counselor_id = ? AND day_schedule = ?", id, day).First(&counselorSchedule)

	if errGetDaySchedule.Error != nil {
		fmt.Errorf(errGetDaySchedule.Error.Error())
		return nil, fmt.Errorf("Schedule not found")
	}

	if errGetDaySchedule.RowsAffected == 0 {
		return nil, fmt.Errorf("Schedule not found")
	}

	return &counselorSchedule, nil

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
