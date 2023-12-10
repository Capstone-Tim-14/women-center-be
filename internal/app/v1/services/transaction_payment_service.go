package services

import (
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	payment "woman-center-be/pkg/payment/midtrans"

	"github.com/midtrans/midtrans-go/coreapi"
	uuid "github.com/satori/go.uuid"
)

type TransactionPaymentService interface {
	ChargeTransactionBooking(orderId uuid.UUID, paymentCode string) error
}

type TransactionPaymentServiceImpl struct {
	BookingRepo     repositories.BookingCounselingRepository
	MidtransService payment.MidtransCoreApi
}

func NewTransactionPaymentService(transactionPayment TransactionPaymentServiceImpl) TransactionPaymentService {
	return &transactionPayment
}

func (service *TransactionPaymentServiceImpl) TransactionPaymentService(orderId uuid.UUID, paymentCode string) (*coreapi.ChargeResponse, error) {

	getBookingData, errGetBooking := service.BookingRepo.FindByOrderId(orderId)

	if errGetBooking != nil {
		return errGetBooking
	}

	RequestMidtransCharge := conversion.ConversionPaymentCharge(*getBookingData, paymentCode)

	CreateChargeMidtrans, errCreateCharge := service.MidtransService.ChargeTransaction(RequestMidtransCharge)

	if errCreateCharge != nil {
		return
	}

}
