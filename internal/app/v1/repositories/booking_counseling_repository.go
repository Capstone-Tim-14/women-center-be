package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BookingCounselingRepository interface {
	CreateBooking(booking *domain.BookingCounseling) (*domain.BookingCounseling, error)
	FindByOrderId(orderId uuid.UUID) (*domain.BookingCounseling, error)
	UpdateStatusBooking(orderId uuid.UUID, status string) (bool, error)
	GetBookingListByCounselor(counselor_id uint) ([]domain.CounselingSession, error)
	GetBookingCounselingDetail(counselor_id uint, user_id uint) (*domain.CounselingSessionDetail, []domain.CounselingScheduleSession, error)
}

type BookingCounselingRepositoryImpl struct {
	db *gorm.DB
}

func NewBookingCounselingRepository(db *gorm.DB) BookingCounselingRepository {
	return &BookingCounselingRepositoryImpl{
		db: db,
	}
}

func (repository *BookingCounselingRepositoryImpl) GetBookingCounselingDetail(counselor_id uint, user_id uint) (*domain.CounselingSessionDetail, []domain.CounselingScheduleSession, error) {

}

func (repository *BookingCounselingRepositoryImpl) GetBookingListByCounselor(counselor_id uint) ([]domain.CounselingSession, error) {
	var CounselingSessionBooked []domain.CounselingSession

	errGetBookingCounseling := repository.db.Raw("SELECT booking.status,booking.user_id,booking.order_id,GROUP_CONCAT(counselor_schedule.time_start) AS time_starts,GROUP_CONCAT(counselor_schedule.time_finish) AS time_finishs,counselor_schedule.day_schedule,user_schedule.date_schedule,users.first_name,users.last_name,package.title AS package_title,counselor_schedule.time_start,counselor_schedule.time_finish FROM booking_counselings booking INNER JOIN booking_counseling_details detail ON detail.id = booking.booking_detail_id INNER JOIN counseling_packages package ON package.id = detail.counseling_package_id INNER JOIN user_schedule_counselings user_schedule ON user_schedule.booking_detail_id = detail.id INNER JOIN users ON user_schedule.user_id = users.id INNER JOIN credentials auth ON auth.id = users.credential_id INNER JOIN counseling_schedules counselor_schedule ON counselor_schedule.id = user_schedule.counselor_schedule_id WHERE counselor_schedule.counselor_id = ? AND booking.status = ? GROUP BY user_schedule.id,booking.order_id,package_title,booking.status,booking.user_id", counselor_id, "SETTLEMENT").Scan(&CounselingSessionBooked)

	if errGetBookingCounseling.Error != nil {
		return nil, fmt.Errorf("Error when find counseling data")
	}

	if errGetBookingCounseling.RowsAffected == 0 {
		return nil, fmt.Errorf("Error when find counseling data")
	}

	return CounselingSessionBooked, nil
}

func (repository *BookingCounselingRepositoryImpl) CreateBooking(booking *domain.BookingCounseling) (*domain.BookingCounseling, error) {
	transaction := repository.db.Begin()

	result := transaction.Create(&booking)
	if result.Error != nil {
		transaction.Rollback()
		return nil, result.Error
	}

	transaction.Commit()

	return booking, nil
}

func (repository *BookingCounselingRepositoryImpl) FindByOrderId(orderId uuid.UUID) (*domain.BookingCounseling, error) {
	var booking *domain.BookingCounseling

	errGetbooking := repository.db.Preload("User").
		Preload("User.Credential").
		Preload("BookingDetail").
		Preload("BookingDetail.Package").
		Preload("BookingDetail.User_Schedules").
		First(&booking, "order_id = ? AND status = ?", orderId, "IN PROCESS")

	if errGetbooking.Error != nil {
		fmt.Errorf(errGetbooking.Error.Error())
		return nil, fmt.Errorf("Error to find booking")
	}

	if errGetbooking.RowsAffected == 0 {
		return nil, fmt.Errorf("Booking not found")
	}

	return booking, nil

}

func (repository *BookingCounselingRepositoryImpl) UpdateStatusBooking(orderId uuid.UUID, status string) (bool, error) {

	var booking *domain.BookingCounseling

	errGetTransaction := repository.db.Preload("User").
		Preload("User.Credential").
		Preload("BookingDetail").
		Preload("BookingDetail.Package").
		Preload("BookingDetail.User_Schedules").
		First(&booking, "order_id = ?", orderId)

	if errGetTransaction.Error != nil {
		return false, errGetTransaction.Error
	}

	errUpdateStatus := repository.db.Model(&booking).Update("status", status)

	if errUpdateStatus.Error != nil {
		return false, fmt.Errorf("Error when update status")
	}

	return true, nil

}
