package domain

import (
	"time"

	"gorm.io/gorm"
)

type Article_Has_Tags struct {
	Id uint
	//Article_id uint
	//Article    *Articles
	Tag_id uint
	Tag    *Tag_Article

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
