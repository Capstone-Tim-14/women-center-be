package domain

import (
	"time"

	"gorm.io/gorm"
)

type Credentials struct {
	Id        uint
	Role_id   uint
	Role      *Roles
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
