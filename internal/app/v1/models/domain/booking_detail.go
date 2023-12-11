package domain

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BookingCounselingDetail struct {
	Id                    uint
	Counseling_package_id uint
	Package               *CounselingPackage `gorm:"foreignKey:Counseling_package_id;references:Id"`
	Tax                   decimal.Decimal
	Total                 decimal.Decimal
	CreatedAt             time.Time
	UpdatedAt             time.Time
	DeletedAt             gorm.DeletedAt
}
