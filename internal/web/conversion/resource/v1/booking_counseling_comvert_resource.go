package conversion

import (
	"strconv"
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

func CounselingSessionBookedConvert(CounselingBooked []domain.CounselingSession) []resources.CounselingSessioningResource {

	var CounselingSessionings []resources.CounselingSessioningResource

	for _, item := range CounselingBooked {
		CounselingSessionings = append(CounselingSessionings, resources.CounselingSessioningResource{
			OrderId:       item.OrderId,
			UserId:        strconv.Itoa(int(item.User_id)),
			FullName:      item.First_name + " " + item.Last_name,
			Package:       item.Package_title,
			Date_schedule: helpers.ParseOnlyDate(helpers.ParseStringToTime(item.Date_schedule)),
			Time_start:    helpers.ParseTimeToClock(helpers.ParseStringToTime(item.Time_start)),
			Time_finisih:  helpers.ParseTimeToClock(helpers.ParseStringToTime(item.Time_finish)),
		})
	}

	return CounselingSessionings

}
