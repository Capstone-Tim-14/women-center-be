package resources

import uuid "github.com/satori/go.uuid"

type BookingCounselingResource struct {
	Order_id          uuid.UUID                   `json:"order_id,omitempty"`
	Transaction_date  string                      `json:"transaction_date,omitempty"`
	Status            string                      `json:"status,omitempty"`
	UserBooking       UserBookingCounseling       `json:"user_booking,omitempty"`
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

type CounselingSessioningResource struct {
	OrderId       string `json:"booking_id,omitempty"`
	UserId        string `json:"user_id,omitempty"`
	FullName      string `json:"fullname,omitempty"`
	Package       string `json:"package,omitempty"`
	Date_schedule string `json:"date_schedule,omitempty"`
	Time_start    string `json:"time_start,omitempty"`
	Time_finisih  string `json:"time_finish,omitempty"`
}

type CounselingSessionDetailResource struct {
	OrderId         string                              `json:"order_id,omitempty"`
	Fulltime        string                              `json:"fulltime,omitempty"`
	Package_title   string                              `json:"package,omitempty"`
	Status          string                              `json:"status,omitempty"`
	Email           string                              `json:"email,omitempty"`
	ScheduleSession []CounselingScheduleSessionResource `json:"schedules,omitempty"`
}

type CounselingScheduleSessionResource struct {
	Day_schedule string `json:"date_schedule,omitempty"`
	Time_start   string `json:"time_start,omitempty"`
	Time_finish  string `json:"time_finish,omitempty"`
}
