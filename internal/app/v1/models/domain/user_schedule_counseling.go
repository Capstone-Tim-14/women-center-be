package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserScheduleCounseling struct {
	Id            uint
	User_id       uint
	User          *Users `gorm:"foreignKey:User_id;references:Id;"`
	Counselor_id  uint
	Counselor     *Counselors `gorm:"foreignKey:Counselor_id;references:Id;"`
	Date_schedule time.Time
	Time_start    string
	Note          string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
