package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *domain.Roles) (*domain.Roles, error)
}

type RoleRepositoryImpl struct {
	db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: db,
	}
}

func (repository *RoleRepositoryImpl) CreateRole(role *domain.Roles) (*domain.Roles, error) {
	result := repository.db.Create(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return role, nil

}
