package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"

	"github.com/gosimple/slug"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func ConversionPaymentCharge(booking domain.BookingCounseling, paymentType string) *coreapi.ChargeReq {

	var PaymentType coreapi.CoreapiPaymentType

	if paymentType == string(midtrans.BankBca) || paymentType == string(midtrans.BankMandiri) || paymentType == string(midtrans.BankBri) || paymentType == string(midtrans.BankPermata) || paymentType == string(midtrans.BankMega) || paymentType == string(midtrans.BankCimb) || paymentType == string(midtrans.BankMaybank) || paymentType == string(midtrans.BankBni) {
		PaymentType = coreapi.PaymentTypeBankTransfer
	} else if paymentType == "gopay" {
		PaymentType = coreapi.PaymentTypeGopay
	} else {
		paymentType = string(coreapi.PaymentTypeQris)
	}

	chargeRequest := &coreapi.ChargeReq{
		PaymentType: PaymentType,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  booking.OrderId.String(),
			GrossAmt: booking.BookingDetail.Total.IntPart(),
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: booking.User.First_name,
			LName: booking.User.Last_name,
			Email: booking.User.Credential.Email,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    slug.Make(booking.BookingDetail.Package.Title),
				Name:  booking.BookingDetail.Package.Title,
				Price: booking.BookingDetail.Package.Price.IntPart(),
				Qty:   1,
			},
		},
	}

	return chargeRequest
}
