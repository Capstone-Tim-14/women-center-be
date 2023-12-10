package domain

import (
	"time"

	"gorm.io/gorm"
)

type Counseling_Schedule struct {
	Id           uint
	Counselor_id uint
	Counselor    *Counselors
	Day_schedule string
	Time_start   time.Time
	Time_finish  time.Time
	Status       string `gorm:"type:enum('OPEN', 'BOOKED', 'CLOSED');default:'OPEN'"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}
