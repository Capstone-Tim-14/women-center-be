package schema

import (
	"time"

	"gorm.io/gorm"
)

type UserScheduleCounseling struct {
	Id                    uint `gorm:"primaryKey"`
	User_id               uint
	Booking_detail_id     uint
	User                  *Users `gorm:"foreignKey:User_id;references:Id;"`
	Counselor_schedule_id uint
	Counselor_schedule    *Counseling_Schedule `gorm:"foreignKey:Counselor_schedule_id;references:Id;"`
	Date_schedule         time.Time
	Time_start            string `gorm:"type:varchar(50)"`
	Note                  string
	CreatedAt             time.Time      `gorm:"autoCreateTime"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt             gorm.DeletedAt `gorm:"index"`
}
