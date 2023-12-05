package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type PaymentMethodRepository interface {
	GetAllBankMethod() ([]domain.BankMethod, error)
}

type PaymentMethodRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentMethodRepository(Db *gorm.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{
		db: Db,
	}
}

func (repository *PaymentMethodRepositoryImpl) GetAllBankMethod() ([]domain.BankMethod, error) {

	var BankMethod []domain.BankMethod

	errGetAll := repository.db.Find(&BankMethod)

	if errGetAll.Error != nil {
		fmt.Errorf(errGetAll.Error.Error())
		return nil, fmt.Errorf("Error to get bank method")
	}

	if errGetAll.RowsAffected == 0 {
		return nil, fmt.Errorf("Bank method is empty")
	}

	return BankMethod, nil
}
