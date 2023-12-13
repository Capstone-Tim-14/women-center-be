package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Event struct {
	Id          uint
	Title       string
	Poster      string
	Locations   string
	Date        *time.Time
	Price       decimal.Decimal
	EventUrl    string
	Time_start  time.Time
	Time_finish time.Time
	EventType   string `gorm:"type:enum('OFFLINE', 'ONLINE');"`
	Status      string `gorm:"type:enum('OPEN', 'CLOSED');default:OPEN"`
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
