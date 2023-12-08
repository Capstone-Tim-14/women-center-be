package services

import (
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/resources/v1"
)

type PaymentMethodService interface {
	GetListBankMethod() ([]resources.BankMethodResource, error)
}

type PaymentMethodServiceImpl struct {
	PaymentMethodRepo repositories.PaymentMethodRepository
}

func NewPaymentMethodService(paymentService PaymentMethodServiceImpl) PaymentMethodService {
	return &paymentService
}

func (service *PaymentMethodServiceImpl) GetListBankMethod() ([]resources.BankMethodResource, error) {

	bankMethodList, errGetList := service.PaymentMethodRepo.GetAllBankMethod()

	if errGetList != nil {
		return nil, errGetList
	}

	ConvertResource := conversion.DomainBankMethodToResource(bankMethodList)

	return ConvertResource, nil

}
