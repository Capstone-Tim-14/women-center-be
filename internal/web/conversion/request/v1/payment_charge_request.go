package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"

	"github.com/gosimple/slug"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

func ConversionPaymentCharge(booking domain.BookingCounseling, paymentType string) *coreapi.ChargeReq {

	var PaymentType coreapi.CoreapiPaymentType
	var paymentBank midtrans.Bank

	if paymentType == string(midtrans.BankBca) {
		paymentBank = midtrans.BankBca
	}

	if paymentType == string(midtrans.BankMandiri) {
		paymentBank = midtrans.BankMandiri
	}

	if paymentType == string(midtrans.BankBri) {
		paymentBank = midtrans.BankBri
	}

	if paymentType == string(midtrans.BankMega) {
		paymentBank = midtrans.BankMega
	}

	if paymentType == string(midtrans.BankCimb) {
		paymentBank = midtrans.BankCimb
	}

	if paymentType == string(midtrans.BankMaybank) {
		paymentBank = midtrans.BankMaybank
	}

	if paymentType == string(midtrans.BankBni) {
		paymentBank = midtrans.BankBni
	}

	if paymentType == string(midtrans.BankPermata) {
		paymentBank = midtrans.BankPermata
	}

	switch paymentType {
	case string(midtrans.BankBca), string(midtrans.BankMandiri), string(midtrans.BankBri), string(midtrans.BankPermata), string(midtrans.BankMega), string(midtrans.BankCimb), string(midtrans.BankMaybank), string(midtrans.BankBni):
		PaymentType = coreapi.PaymentTypeBankTransfer
	case "gopay":
		PaymentType = coreapi.PaymentTypeGopay
	default:
		PaymentType = coreapi.PaymentTypeQris
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
				Price: booking.BookingDetail.Total.IntPart(),
				Qty:   1,
			},
		},
	}

	if PaymentType == coreapi.PaymentTypeBankTransfer {
		chargeRequest.BankTransfer = &coreapi.BankTransferDetails{
			Bank: paymentBank,
		}
	}

	return chargeRequest
}
