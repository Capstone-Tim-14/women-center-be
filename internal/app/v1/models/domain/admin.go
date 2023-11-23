package domain

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id            uint
	Credential_id uint
	Credential    *Credentials
	First_name    string
	Last_name     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
