package schema

import (
	"time"

	"gorm.io/gorm"
)

type Articles struct {
	Id            uint   `gorm:"primaryKey;"`
	Title         string `gorm:"type:varchar(100)"`
	Slug          string `gorm:"type:varchar(100)"`
	Content       string
	Status        string         `gorm:"type:varchar(10);default:REVIEW"`
	PublishedAt   time.Time      `gorm:"autoCreateTime"`
	Admin_id      *uint          `gorm:"default:null"`
	Admin         *Admin         `gorm:"foreignKey:Admin_id;"`
	Counselors_id *uint          `gorm:"default:null"`
	Counselors    *Counselors    `gorm:"foreignKey:Counselors_id;"`
	Thumbnail     string         `gorm:"type:varchar(255)"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
