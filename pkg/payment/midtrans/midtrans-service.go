package payment

import (
	"fmt"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/spf13/viper"
)

type MidtransCoreApi interface {
	ChargeTransaction(request *coreapi.ChargeReq) (*coreapi.ChargeResponse, error)
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
