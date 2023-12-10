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
}

type BookingCounselingRepositoryImpl struct {
	db *gorm.DB
}

func NewBookingCounselingRepository(db *gorm.DB) BookingCounselingRepository {
	return &BookingCounselingRepositoryImpl{
		db: db,
	}
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
		First(&booking, "order_id = ?", orderId)

	if errGetbooking.Error != nil {
		fmt.Errorf(errGetbooking.Error.Error())
		return nil, fmt.Errorf("Error to find booking")
	}

	if errGetbooking.RowsAffected == 0 {
		return nil, fmt.Errorf("Booking not found")
	}

	return booking, nil

}
