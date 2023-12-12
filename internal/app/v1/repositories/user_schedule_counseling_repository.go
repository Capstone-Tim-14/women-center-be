package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"

	"gorm.io/gorm"
)

type UserScheduleCounselingRepository interface {
	FindScheduleByDateAndTimeExist(request requests.UserScheduleCounselingQueryRequest) (bool, error)
	CreateUserScheduling(user domain.Users, request domain.UserScheduleCounseling) (*domain.UserScheduleCounseling, error)
	UpdateMultipleScheduleBooked(schedules []domain.UserScheduleCounseling, booking_detail_id uint) error
}

type UserScheduleCounselingRepositoryImpl struct {
	Db *gorm.DB
}

func NewUserScheduleConselingRepository(db *gorm.DB) UserScheduleCounselingRepository {
	return &UserScheduleCounselingRepositoryImpl{
		Db: db,
	}
}

func (repo *UserScheduleCounselingRepositoryImpl) UpdateMultipleScheduleBooked(schedules []domain.UserScheduleCounseling, booking_detail_id uint) error {

	transaction := repo.Db.Begin()

	for _, schedule := range schedules {
		errUpdateBookingDetail := transaction.Model(&schedule).Update("booking_detail_id", booking_detail_id)

		if errUpdateBookingDetail.Error != nil {
			transaction.Rollback()
			fmt.Errorf(errUpdateBookingDetail.Error.Error())
			return fmt.Errorf("Error update booking detail")
		}
	}

	transaction.Commit()

	return nil
}

func (repo *UserScheduleCounselingRepositoryImpl) FindScheduleByDateAndTimeExist(request requests.UserScheduleCounselingQueryRequest) (bool, error) {

	var Schedule *domain.UserScheduleCounseling

	ErrSchedule := repo.Db.Where("counselor_schedule_id = ? AND date_schedule = ? AND time_start = ?", request.Counselor_schedule_id, request.Day, request.Time_start).First(&Schedule)

	if ErrSchedule.Error != nil {
		return false, ErrSchedule.Error
	}

	return true, nil
}

func (repo *UserScheduleCounselingRepositoryImpl) CreateUserScheduling(user domain.Users, request domain.UserScheduleCounseling) (*domain.UserScheduleCounseling, error) {

	var CounselingSchedule domain.Counseling_Schedule

	transaction := repo.Db.Begin()

	if GetCounselingSchedule := transaction.First(&CounselingSchedule, "id = ?", request.Counselor_schedule_id); GetCounselingSchedule.Error != nil {
		transaction.Rollback()
		return nil, fmt.Errorf("Counseling schedule not found")
	}

	if ErrCreateSchedule := transaction.Model(&user).Association("UserScheduleCounseling").Append(&request); ErrCreateSchedule != nil {
		transaction.Rollback()
		return nil, fmt.Errorf("Error create schedule user")
	}

	transaction.Commit()

	return &request, nil

}
