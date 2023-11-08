package domain

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	Id        uint
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
