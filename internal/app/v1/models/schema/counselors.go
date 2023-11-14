package schema

import (
	"time"

	"gorm.io/gorm"
)

type Counselors struct {
	Id                uint `gorm:"primaryKey;"`
	Credential_id     uint
	Credential        *Credentials `gorm:"foreignKey:Credential_id;references:Id;"`
	First_name        string       `gorm:"type:varchar(100)"`
	Last_name         string       `gorm:"type:varchar(100)"`
	Email             string       `gorm:"type:varchar(255)"`
	Profile_picture   string       `gorm:"varchar(255)"`
	Phone_number      string       `gorm:"varchar(255)"`
	Description       string
	Address           string
	Certification_url string         `gorm:"type:varchar(255)"`
	Status            string         `gorm:"type:varchar(10);default:INACTIVE"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}