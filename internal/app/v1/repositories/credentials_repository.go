package repositories

import (
	"woman-center-be/internal/app/v1/models/domain"

	"gorm.io/gorm"
)

type CredentialRepository interface {
	CheckEmailCredential(email string) (*domain.Credentials, error)
	GetAuthUser(id uint, role string) (*domain.Users, *domain.Counselors, error)
}

type CredentialRepositoryImpl struct {
	Db *gorm.DB
}

func NewCredentialRepository(db *gorm.DB) CredentialRepository {
	return &CredentialRepositoryImpl{
		Db: db,
	}
}

func (c *CredentialRepositoryImpl) CheckEmailCredential(email string) (*domain.Credentials, error) {

	var Credential *domain.Credentials

	if ErrCheckEmail := c.Db.Preload("Role").First(&Credential, "email = ?", email).Error; ErrCheckEmail != nil {
		return nil, ErrCheckEmail
	}

	return Credential, nil
}

func (c *CredentialRepositoryImpl) GetAuthUser(id uint, role string) (*domain.Users, *domain.Counselors, error) {

	var user *domain.Users
	var counselor *domain.Counselors
	var ErrGetUserData error

	if role == "user" {
		if ErrGetUserData = c.Db.Preload("Credential").Preload("Credential.Role").First(&user, "Credential_id = ?", id).Error; ErrGetUserData == nil {
			return user, nil, nil
		}

	} else if role == "counselor" {

		if ErrGetUserData = c.Db.Preload("Credential").Preload("Credential.Role").First(&counselor, "Credential_id = ?", id).Error; ErrGetUserData == nil {
			return nil, counselor, nil
		}

	}

	return nil, nil, ErrGetUserData

}
