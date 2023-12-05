package schema

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type CounselingPackage struct {
	Id               uint   `gorm:"primaryKey;"`
	Title            string `gorm:"type:varchar(100)"`
	Thumbnail        string `gorm:"type:varchar(255)"`
	Session_per_week uint
	Price            decimal.Decimal `gorm:"type:decimal(10,2);"`
	Description      string
	CreatedAt        time.Time      `gorm:"autoCreateTime"`
	UpdatedAt        time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}
