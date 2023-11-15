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
	Status        string    `gorm:"type:varchar(10);default:Need To Review"`
	PublishedAt   time.Time `gorm:"type:datetime"`
	Admin_id      uint
	Admin         *Admin `gorm:"foreignKey:Admin_id;references:Id;"`
	Counselors_id uint
	Counselors    *Counselors    `gorm:"foreignKey:Counselors_id;references:Id;"`
	Thumbnail     string         `gorm:"type:varchar(100)"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
