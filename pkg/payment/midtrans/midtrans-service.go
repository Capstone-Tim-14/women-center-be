package payment

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/spf13/viper"
)

type MidtransCoreApi interface {
	ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
	CheckTransactionPayment(transactionData map[string]interface{}) (string, error)
}

type MidtransCoreApiImpl struct {
	CoreApi coreapi.Client
}

func NewMidtransCoreApiImpl() MidtransCoreApi {

	core := coreapi.Client{}

	core.New(viper.GetString("MIDTRANS.SERVER_KEY"), midtrans.Sandbox)

	return &MidtransCoreApiImpl{
		CoreApi: core,
	}
}

func (payment *MidtransCoreApiImpl) ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error) {

	response, errRes := payment.CoreApi.ChargeTransaction(request)

	if errRes != nil {
		fmt.Errorf(errRes.Error())
		return nil, fmt.Errorf("Error create transaction")
	}

	return response, nil

}

func (payment *MidtransCoreApiImpl) CheckTransactionPayment(orderId string) (string, error) {

	transactionStatusReps, errGetTransaction := payment.CoreApi.CheckTransaction(orderId)

	if errGetTransaction != nil {
		return "", fmt.Errorf("Transaction not found")
	}

	if transactionStatusReps != nil {
		if transactionStatusReps.TransactionStatus == "capture" {
			if transactionStatusReps.FraudStatus == "challenge" {
				return transactionStatusReps.FraudStatus, nil
			} else if transactionStatusReps.FraudStatus == "accept" {
				return transactionStatusReps.FraudStatus, nil
			}
		} else {
			return transactionStatusReps.TransactionStatus, nil
		}
	}

	return "", fmt.Errorf("Transaction invalid")
}
