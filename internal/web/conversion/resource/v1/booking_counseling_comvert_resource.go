package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func BookingCounselingToDomainBookingCounseling(transaction *domain.BookingCounseling, counselingPackage *domain.CounselingPackage, user *domain.Users, Counselor *domain.Counselors, schedule []domain.UserScheduleCounseling) *resources.BookingCounselingResource {
	result := resources.BookingCounselingResource{
		Order_id:         transaction.OrderId,
		Transaction_date: helpers.ParseOnlyDate(&transaction.Transaction_date),
		Status:           transaction.Status,
		UserBooking: resources.UserBookingCounseling{
			FullName: user.First_name + " " + user.Last_name,
			Email:    user.Credential.Email,
		},
		Counselor: Counselor.First_name + " " + Counselor.Last_name,
		CounselingPackege: resources.CouselingPackage{
			Title: counselingPackage.Title,
			Price: int(counselingPackage.Price.IntPart()),
		},
	}
	for _, schedule := range schedule {
		result.Schedule = append(result.Schedule, resources.BookingCounselingSchedule{
			Date: helpers.ParseOnlyDate(&schedule.Date_schedule),
			Time: schedule.Time_start,
		})
	}
	result.Tax = int(transaction.BookingDetail.Tax.IntPart())
	result.Total = int(transaction.BookingDetail.Total.IntPart())

	return &result
}
