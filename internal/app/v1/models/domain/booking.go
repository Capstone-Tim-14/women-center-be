package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BookingCounseling struct {
	Id                    uint
	OrderId               uuid.UUID
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

func (booking *BookingCounseling) BeforeCreate(tx *gorm.DB) (err error) {
	booking.OrderId = uuid.NewV4()

	return
}
