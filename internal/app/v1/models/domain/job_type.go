package domain

import (
	"time"

	"gorm.io/gorm"
)

type Job_Type struct {
	Id          uint
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}
