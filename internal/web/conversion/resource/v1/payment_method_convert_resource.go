package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func DomainBankMethodToResource(bankMethod []domain.BankMethod) []resources.BankMethodResource {

	var BankMethodResources []resources.BankMethodResource

	for _, bank := range bankMethod {
		BankMethodResources = append(BankMethodResources, resources.BankMethodResource{
			Title:    bank.Title,
			BankCode: bank.BankCode,
			Image:    bank.Image,
		})
	}

	return BankMethodResources

}
