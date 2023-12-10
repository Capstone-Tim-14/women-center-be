package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
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
