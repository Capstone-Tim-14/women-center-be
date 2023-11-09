package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CreateAdmin(admin *domain.Admin) (*domain.Admin, error)
}

type AdminRepositoryImpl struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &AdminRepositoryImpl{
		db: db,
	}
}

func (repository *AdminRepositoryImpl) CreateAdmin(admin *domain.Admin) (*domain.Admin, error) {
	result := admin.db.Create(&admin)
	if result.Error != nil {
		return nil, result.Error
	}

	return admin, nil
}
