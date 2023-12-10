package schema

import (
	"time"

	"gorm.io/gorm"
)

type Counseling_Schedule struct {
	Id           uint `gorm:"primaryKey;"`
	Counselor_id uint
	Counselor    *Counselors    `gorm:"foreignKey:Counselor_id;references:Id;"`
	Day_schedule string         `gorm:"type:varchar(10)"`
	Time_start   time.Time      `gorm:"type:time"`
	Time_finish  time.Time      `gorm:"type:time"`
	Status       string         `gorm:"type:enum('OPEN', 'BOOKED', 'CLOSED');default:'OPEN'"`
	CreatedAt    time.Time      `gorm:"autoCreateTime"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
