package resources

type BookingCounselingResource struct {
	Booking_id        uint                        `json:"booking_id,omitempty"`
	Transaction_date  string                      `json:"transaction_date,omitempty"`
	Status            string                      `json:"status,omitempty"`
	UserBooking       UserBookingCounseling       `json:"use_booking,omitempty"`
	Counselor         string                      `json:"counselor,omitempty"`
	CounselingPackege CouselingPackage            `json:"counseling_package,omitempty"`
	Schedule          []BookingCounselingSchedule `json:"counseling_schedule,omitempty"`
	Tax               int                         `json:"tax,omitempty"`
	Total             int                         `json:"total,omitempty"`
}

type UserBookingCounseling struct {
	FullName string `json:"full_name,omitempty"`
	Email    string `json:"email,omitempty"`
}

type CouselingPackage struct {
	Title string `json:"title,omitempty"`
	Price int    `json:"price,omitempty"`
}

type BookingCounselingSchedule struct {
	Date string `json:"date,omitempty"`
	Time string `json:"time,omitempty"`
}
