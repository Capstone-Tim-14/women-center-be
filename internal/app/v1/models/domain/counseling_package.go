package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CounselingPackage struct {
	Id               uint
	Title            string
	Thumbnail        *string
	Session_per_week uint
	Price            decimal.Decimal
	Description      string
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}
