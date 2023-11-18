package schema

import (
	"time"

	"gorm.io/gorm"
)

type Article_Has_Tags struct {
	Id uint `gorm:"primaryKey;"`
	//Article_id uint
	//Article    *Articles `gorm:"foreignKey:Article_id;references:Id;"`
	Tag_id uint
	Tag    *Tag_Article `gorm:"foreignKey:Tag_id;references:Id;"`

	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
