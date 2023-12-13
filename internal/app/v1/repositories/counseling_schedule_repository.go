package repositories

import (
	"fmt"
	"time"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateSchedule(counselor *domain.Counselors, scheduling []domain.Counseling_Schedule) error
	CheckDayCounselingScheduleExists(id int, day string) (*domain.Counseling_Schedule, error)
	FindStartEndDateCounseling(counselor_id int, day string, start time.Time, finish time.Time) (*domain.Counseling_Schedule, error)
	GroupingStartTimeAndFinishTimeCounseling(counselor_id int) ([]domain.Counseling_Schedule, error)
}

type ScheduleRepositoryImpl struct {
	Db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &ScheduleRepositoryImpl{
		Db: db,
	}
}
func (repository *ScheduleRepositoryImpl) GroupingStartTimeAndFinishTimeCounseling(counselor_id int) ([]domain.Counseling_Schedule, error) {

	var counselorSchedules []domain.Counseling_Schedule

	GetListCounselor := repository.Db.Select("day_schedule,GROUP_CONCAT(time_start) AS time_starts,GROUP_CONCAT(time_finish) AS time_finishs").
		Where("counselor_id = ?", counselor_id).
		Group("day_schedule").
		Find(&counselorSchedules)

	if GetListCounselor.Error != nil {
		return nil, fmt.Errorf("Error to get schedules")
	}

	if GetListCounselor.RowsAffected == 0 {
		return nil, fmt.Errorf("schedules not found")
	}

	return counselorSchedules, nil

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

func (repository *ScheduleRepositoryImpl) FindStartEndDateCounseling(counselor_id int, day string, start time.Time, finish time.Time) (*domain.Counseling_Schedule, error) {
	var Schedule domain.Counseling_Schedule

	GetScheduleData := repository.Db.Where("counselor_id = ? AND day_schedule = ? AND time_start = ? AND time_finish = ?", counselor_id, day, start, finish).First(&Schedule)
	if GetScheduleData.Error != nil {
		return nil, GetScheduleData.Error
	}

	return &Schedule, nil
}
