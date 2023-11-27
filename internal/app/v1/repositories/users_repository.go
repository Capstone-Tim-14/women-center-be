package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *domain.Users) (*domain.Users, error)
	FindyByEmail(email string) (*domain.Users, error)
	FindByID(id int) (*domain.Users, error)
	UpdateUser(user *domain.Users) (*domain.Users, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
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

	result := repository.db.Preload("Credential").Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (repository *UserRepositoryImpl) UpdateUser(user *domain.Users) (*domain.Users, error) {
	result := repository.db.Save(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
