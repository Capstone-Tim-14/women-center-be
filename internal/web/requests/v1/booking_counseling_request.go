package requests

import (
	"time"

	"github.com/shopspring/decimal"
)

type BookingCounselingRequest struct {
	Booking_date string `json:"booking_date" validate:"required"`
	Booking_time string `json:"booking_time" validate:"required"`
}

type TransactionBookingRequest struct {
	User_id            uint
	Transaction_date   time.Time
	Status             string
	Transaction_detail TransactionBookingDetailRequest
}

type TransactionBookingDetailRequest struct {
	Counseling_package_id uint
	Tax                   decimal.Decimal
	Total                 decimal.Decimal
}
