package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/helpers"
)

func UserScheduleCounselingToDomainSchedule(request requests.BookingCounselingRequest, user_id uint, schedule_id uint) *domain.UserScheduleCounseling {
	return &domain.UserScheduleCounseling{
		Date_schedule:         *helpers.ParseStringToTime(request.Booking_date),
		Time_start:            request.Booking_time,
		Counselor_schedule_id: schedule_id,
		User_id:               user_id,
	}
}
