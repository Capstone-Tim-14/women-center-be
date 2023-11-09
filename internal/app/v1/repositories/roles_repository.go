package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type RoleRepository interface {
	CreateRole(role *domain.Roles) (*domain.Roles, error)
	FindById(id int) (*domain.Roles, error)
	FindByName(name string) (*domain.Roles, error)
	FindAll() ([]domain.Roles, error)
	DeleteById(id int) error
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

func (repository *RoleRepositoryImpl) FindById(id int) (*domain.Roles, error) {
	role := domain.Roles{}

	result := repository.db.Where("id = ?", id).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (repository *RoleRepositoryImpl) FindByName(name string) (*domain.Roles, error) {
	role := domain.Roles{}

	result := repository.db.Where("name = ?", name).First(&role)
	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (repository *RoleRepositoryImpl) FindAll() ([]domain.Roles, error) {
	roles := []domain.Roles{}

	result := repository.db.Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}

	return roles, nil
}

func (repository *RoleRepositoryImpl) DeleteById(id int) error {
	result := repository.db.Table("roles").Where("id = ?", id).Unscoped().Delete(id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
