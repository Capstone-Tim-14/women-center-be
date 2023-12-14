package domain

import (
	"time"

	"gorm.io/gorm"
)

type Counselors struct {
	Id              uint
	Credential_id   uint
	Credential      *Credentials
	First_name      string
	Last_name       string
	Profile_picture string `gorm:"default:https://pub-86c5755f32914550adb162dd2b8850d0.r2.dev/default-profile.jpg"`
	Description     string
	Education       string
	Birthday        *time.Time
	Status          string
	Specialists     []Specialist          `gorm:"many2many:counselor_has_specialists;foreignKey:Id;references:Id;"`
	Schedules       []Counseling_Schedule `gorm:"foreignKey:Counselor_id;references:Id;"`
	SingleSchedule  *Counseling_Schedule  `gorm:"foreignKey:Counselor_id;references:Id;"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	DeletedAt       gorm.DeletedAt
}
