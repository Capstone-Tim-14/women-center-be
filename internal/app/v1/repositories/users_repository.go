package repositories

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.Users) (*domain.Users, error)
	FindyByEmail(email string) (*domain.Users, error)
	FindByID(id int) (*domain.Users, error)
	UpdateUser(user *domain.Users) (*domain.Users, error)
	UpdateOTP(user *domain.Users, secret string) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) UpdateOTP(user *domain.Users, secret string) error {

	var UpdateUserColumn map[string]interface{}

	if secret != "" {
		UpdateUserColumn = map[string]interface{}{
			"secret_otp": secret,
			"otp_enable": true,
		}
	} else {
		UpdateUserColumn = map[string]interface{}{
			"secret_otp": nil,
			"otp_enable": true,
		}
	}

	ErrUpdateOTP := repository.db.Model(&user).Updates(UpdateUserColumn)

	if ErrUpdateOTP.Error != nil {
		fmt.Errorf(ErrUpdateOTP.Error.Error())
		return fmt.Errorf("Error to set opt")
	}

	return nil

}

func (repository *UserRepositoryImpl) CreateUser(user *domain.Users) (*domain.Users, error) {
	result := repository.db.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil

}

func (repository *UserRepositoryImpl) FindyByEmail(email string) (*domain.Users, error) {
	user := domain.Users{}

	result := repository.db.InnerJoins("Credential").InnerJoins("Credential.Role").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) FindByID(id int) (*domain.Users, error) {
	user := domain.Users{}

	result := repository.db.Preload("Credential").Preload("ArticleFavorites").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) UpdateUser(user *domain.Users) (*domain.Users, error) {

	Transaction := repository.db.Begin()

	result := Transaction.Model(&user).Updates(&user)

	if result.Error != nil {
		Transaction.Rollback()
		return nil, result.Error
	}

	result = Transaction.Model(&user.Credential).Updates(&user.Credential)

	if result.Error != nil {
		Transaction.Rollback()
		return nil, result.Error
	}

	Transaction.Commit()

	return user, nil
}
