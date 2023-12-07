package domain

import (
	"time"

	"gorm.io/gorm"
)

type BookingCounseling struct {
	Id                    uint
	User_id               uint
	User                  *Users `gorm:"foreignKey:User_id;references:Id"`
	Booking_counseling_id uint
	BookingDetail         *BookingCounselingDetail `gorm:"foreignKey:Booking_counseling_id;references:Id"`
	Transaction_date      time.Time
	Status                string `gorm:"type:enum('IN PROCESS', 'PENDING', 'FAILED', 'SETTLEMENT');default:IN PROCESS"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
