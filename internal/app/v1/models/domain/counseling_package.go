package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type CounselingPackage struct {
	Id                 uint
	Package_name       string
	Description        string
	PublishedAt        time.Time
	Thumbnail          *string
	Number_of_sessions uint
	Price              decimal.Decimal
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          time.Time
}
