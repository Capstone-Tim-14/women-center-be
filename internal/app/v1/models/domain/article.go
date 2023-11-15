package domain

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	Id            uint
	Title         string
	Slug          string
	Content       string
	Status        string
	PublishedAt   time.Time
	Admin_id      uint
	Admin         *Admin
	Counselors_id uint
	Counselors    *Counselors
	Thumbnail     string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
