package schema

import (
	"time"

	"gorm.io/gorm"
)

type GenderType string

const (
	Male    GenderType = "Male"
	Female  GenderType = "Female"
	Unknown GenderType = "Unknown"
)

type Counselors struct {
	Id            uint `gorm:"primaryKey;"`
	Credential_id uint
	Credential    *Credentials `gorm:"foreignKey:Credential_id;references:Id;"`
	First_name    string       `gorm:"type:varchar(100)"`
	Last_name     string       `gorm:"type:varchar(100)"`
	Email         string       `gorm:"type:varchar(255)"`
	Gender        string       `gorm:"type:enum('Male','Female');default:'Unknown'"`
	Certification string       `gorm:"type:varchar(255)"`
	Description   string
	Social_media  string `gorm:"type:varchar(255)"`
	Status        string `gorm:"type:varchar(10);default:INACTIVE"`
	//Profile_picture string `gorm:"varchar(255)"`
	//Phone_number    int
	//Address         string
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
