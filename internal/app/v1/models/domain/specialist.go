package domain

import (
	"time"

	"gorm.io/gorm"
)

type Specialist struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
