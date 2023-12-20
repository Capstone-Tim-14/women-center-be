package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"
)

func BookingCounselingRequestToDomain(request requests.TransactionBookingRequest) domain.BookingCounseling {
	return domain.BookingCounseling{
		User_id:          request.User_id,
		Transaction_date: request.Transaction_date,
		Status:           request.Status,
		BookingDetail: &domain.BookingCounselingDetail{
			Counseling_package_id: request.Transaction_detail.Counseling_package_id,
			Tax:                   request.Transaction_detail.Tax,
			Total:                 request.Transaction_detail.Total,
		},
	}

}

func ConvertBookingDataRequest(booking domain.BookingCounseling) resources.BookingEmailRequest {

	request := resources.BookingEmailRequest{
		Order_id:         booking.OrderId.String(),
		Transaction_date: helpers.ParseOnlyDate(&booking.Transaction_date),
		Name_customer:    booking.User.First_name + " " + booking.User.Last_name,
		Package_name:     booking.BookingDetail.Package.Title,
		Total:            int(booking.BookingDetail.Total.IntPart()),
	}

	return request

}
