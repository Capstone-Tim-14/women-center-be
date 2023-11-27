package domain

import (
	"time"

	"gorm.io/gorm"
)

type Counselors struct {
	Id                uint
	Credential_id     uint
	Credential        *Credentials
	First_name        string
	Last_name         string
	Profile_picture   string
	Phone_number      string
	Description       string
	Certification_url string
	Status            string
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt
}
