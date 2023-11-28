package schema

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id              uint `gorm:"primaryKey;"`
	Credential_id   uint
	Credential      *Credentials `gorm:"foreignKey:Credential_id;references:Id;"`
	First_name      string       `gorm:"type:varchar(100)"`
	Last_name       string       `gorm:"type:varchar(100)"`
	Profile_picture string       `gorm:"varchar(255)"`
	Phone_number    string       `gorm:"type:varchar(20)"`
	Birthday        *time.Time
	Status          string         `gorm:"type:varchar(10);default:INACTIVE"`
	Secret_Otp      *string        `gorm:"type:varchar(100)"`
	Otp_enable      bool           `gorm:"default:false"`
	CreatedAt       time.Time      `gorm:"autoCreateTime"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
}
