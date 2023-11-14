package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id              int
	Credential_id   uint
	Credential      *Credentials
	First_name      string
	Last_name       string
	Email           string
	Password        string
	Profile_picture string
	Phone_number    string
	Address         string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
