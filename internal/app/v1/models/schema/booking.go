package schema

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type BookingCounseling struct {
	Id                uint      `gorm:"primaryKey"`
	OrderId           uuid.UUID `gorm:"type:char(36);"`
	User_id           uint
	User              *Users `gorm:"foreignKey:User_id;references:Id"`
	Booking_detail_id uint
	BookingDetail     *BookingCounselingDetail `gorm:"foreignKey:Booking_detail_id;references:Id"`
	Transaction_date  time.Time
	Status            string         `gorm:"type:enum('IN PROCESS', 'PENDING', 'FAILED', 'SETTLEMENT');default:IN PROCESS"`
	CreatedAt         time.Time      `gorm:"autoCreateTime"`
	UpdatedAt         time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}
