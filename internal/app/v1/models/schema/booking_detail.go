package schema

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BookingCounselingDetail struct {
	Id                    uint `gorm:"primaryKey"`
	Counseling_package_id uint
	Package               CounselingPackage `gorm:"foreignKey:Counseling_package_id;references:Id"`
	Tax                   decimal.Decimal   `gorm:"type:decimal(10,2)"`
	Total                 decimal.Decimal   `gorm:"type:decimal(10,2)"`
	CreatedAt             time.Time         `gorm:"autoCreateTime"`
	UpdatedAt             time.Time         `gorm:"autoUpdateTime:milli"`
	DeletedAt             gorm.DeletedAt    `gorm:"index"`
}
