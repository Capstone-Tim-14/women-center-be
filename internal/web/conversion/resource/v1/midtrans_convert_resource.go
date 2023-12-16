package conversion

import (
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/utils/helpers"

	"github.com/midtrans/midtrans-go/coreapi"
)

func MidtransConvertToResponse(Midtrans *coreapi.ChargeResponse) resources.MidtransInvoice {

	Response := resources.MidtransInvoice{
		StatusMessage:     Midtrans.StatusMessage,
		TransactionId:     Midtrans.TransactionID,
		OrderId:           Midtrans.OrderID,
		GrossAmount:       Midtrans.GrossAmount,
		PaymentType:       Midtrans.PaymentType,
		TransactionTime:   helpers.ParseOnlyDate(helpers.ParseStringToTime(Midtrans.TransactionTime)),
		Currency:          Midtrans.Currency,
		TransactionStatus: Midtrans.TransactionStatus,
		FraudStatus:       Midtrans.FraudStatus,
		ExpiryTime:        helpers.ParseOnlyDate(helpers.ParseStringToTime(Midtrans.ExpiryTime)),
	}

	for _, va_number := range Midtrans.VaNumbers {
		VaNumber := resources.MidtransVaNumbers{
			Bank:     va_number.Bank,
			VaNumber: va_number.VANumber,
		}

		Response.VaNumbers = append(Response.VaNumbers, VaNumber)
	}

	for _, action := range Midtrans.Actions {
		GetAction := resources.MidtransActions{
			Name:   action.Name,
			Method: action.Method,
			Url:    action.URL,
		}

		Response.Actions = append(Response.Actions, GetAction)
	}

	return Response
}
