package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BookingCounselingDetail struct {
	Id                    uint
	Counseling_package_id uint
	Package               *CounselingPackage
	Tax                   decimal.Decimal
	Total                 decimal.Decimal
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
