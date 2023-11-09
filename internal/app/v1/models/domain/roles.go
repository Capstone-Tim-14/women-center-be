package domain

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
