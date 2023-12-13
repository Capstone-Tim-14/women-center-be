package schema

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Event struct {
	Id          uint            `gorm:"primaryKey"`
	Title       string          `gorm:"type:varchar(255)"`
	Poster      string          `gorm:"type:varchar(255)"`
	Locations   string          `gorm:"type:varchar(50)"`
	Date        time.Time       `gorm:"autoCreateTime"`
	Price       decimal.Decimal `gorm:"type:decimal(10,2)"`
	Time_start  time.Time       `gorm:"type:time"`
	Time_finish time.Time       `gorm:"type:time"`
	EventUrl    string          `gorm:"type:varchar(255)"`
	EventType   string          `gorm:"type:enum('OFFLINE', 'ONLINE');"`
	Status      string          `gorm:"type:enum('OPEN', 'CLOSED');default:OPEN"`
	Description string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
