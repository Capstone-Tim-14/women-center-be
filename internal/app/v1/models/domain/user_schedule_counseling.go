package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserScheduleCounseling struct {
	Id                    uint
	User_id               uint
	User                  *Users `gorm:"foreignKey:User_id;references:Id;"`
	Counselor_schedule_id uint
	Counselor_schedule    *Counseling_Schedule `gorm:"foreignKey:Counselor_schedule_id;references:Id;"`
	Date_schedule         time.Time
	Time_start            string
	Note                  string
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
