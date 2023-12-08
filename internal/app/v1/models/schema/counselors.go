package schema

import (
	"time"

	"gorm.io/gorm"
)

type Counselors struct {
	Id              uint `gorm:"primaryKey;"`
	Credential_id   uint
	Credential      *Credentials `gorm:"foreignKey:Credential_id;references:Id;"`
	First_name      string       `gorm:"type:varchar(100)"`
	Last_name       string       `gorm:"type:varchar(100)"`
	Profile_picture string       `gorm:"varchar(255);default:https://pub-86c5755f32914550adb162dd2b8850d0.r2.dev/default-profile.jpg"`
	Description     string
	Status          string                `gorm:"type:varchar(10);default:INACTIVE"`
	Specialists     []Specialist          `gorm:"many2many:counselor_has_specialists;foreignKey:Id;references:Id;"`
	Schedules       []Counseling_Schedule `gorm:"foreignKey:Counselor_id;references:Id;"`
	CreatedAt       time.Time             `gorm:"autoCreateTime"`
	UpdatedAt       time.Time             `gorm:"autoUpdateTime:mili"`
	DeletedAt       gorm.DeletedAt        `gorm:"index"`
}
