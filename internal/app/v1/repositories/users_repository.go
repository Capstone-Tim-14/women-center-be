package repositories

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser()
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (user *UserRepositoryImpl) CreateUser() {

}
