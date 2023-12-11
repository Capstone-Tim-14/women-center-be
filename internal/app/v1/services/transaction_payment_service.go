package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	Rescon "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/resources/v1"
	payment "woman-center-be/pkg/payment/midtrans"

	uuid "github.com/satori/go.uuid"
)

type TransactionPaymentService interface {
	TransactionPaymentService(orderId uuid.UUID, paymentCode string) (*resources.MidtransInvoice, error)
}

type TransactionPaymentServiceImpl struct {
	BookingRepo     repositories.BookingCounselingRepository
	MidtransService payment.MidtransCoreApi
}

func NewTransactionPaymentService(transactionPayment TransactionPaymentServiceImpl) TransactionPaymentService {
	return &transactionPayment
}

func (service *TransactionPaymentServiceImpl) TransactionPaymentService(orderId uuid.UUID, paymentCode string) (*resources.MidtransInvoice, error) {

	getBookingData, errGetBooking := service.BookingRepo.FindByOrderId(orderId)

	if errGetBooking != nil {
		return nil, errGetBooking
	}

	RequestMidtransCharge := conversion.ConversionPaymentCharge(*getBookingData, paymentCode)

	CreateChargeMidtrans, errCreateCharge := service.MidtransService.ChargeTransaction(RequestMidtransCharge)

	if errCreateCharge != nil {
		return nil, errCreateCharge
	}

	StatusUpdateBooking, errUpdateBooking := service.BookingRepo.UpdateStatusBooking(orderId, "PENDING")

	if !StatusUpdateBooking {
		fmt.Errorf(errUpdateBooking.Error())
		return nil, fmt.Errorf("Error update status transaction")
	}

	ResponseCharge := Rescon.MidtransConvertToResponse(CreateChargeMidtrans)

	return &ResponseCharge, nil

}
