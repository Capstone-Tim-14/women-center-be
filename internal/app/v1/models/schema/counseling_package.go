package schema

import (
	"time"

	"github.com/shopspring/decimal"
)

type CounselingPackage struct {
	Id                 uint   `gorm:"primaryKey;"`
	Package_name       string `gorm:"type:varchar(100)"`
	Description        string
	PublishedAt        time.Time `gorm:"autoCreateTime"`
	Thumbnail          string    `gorm:"type:varchar(255)"`
	Number_of_sessions uint
	Price              decimal.Decimal `gorm:"type:decimal(10,2);"`
	CreatedAt          time.Time       `gorm:"autoCreateTime"`
	UpdatedAt          time.Time       `gorm:"autoCreateTime"`
	DeletedAt          time.Time       `gorm:"autoCreateTime"`
}
