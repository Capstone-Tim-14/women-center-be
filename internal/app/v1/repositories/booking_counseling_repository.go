package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type BookingCounselingRepository interface {
	CreateBooking(booking *domain.BookingCounseling) (*domain.BookingCounseling, error)
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
