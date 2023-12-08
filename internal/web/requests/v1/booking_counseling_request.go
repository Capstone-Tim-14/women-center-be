package requests

type BookingCounselingRequest struct {
	Booking_date string `json:"booking_date" validate:"required"`
	Booking_time string `json:"booking_time" validate:"required"`
}
