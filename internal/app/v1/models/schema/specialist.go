package schema

import (
	"time"

	"gorm.io/gorm"
)

type Specialist struct {
	Id          uint   `gorm:"primaryKey;"`
	Name        string `gorm:"varchar(100)"`
	Description string
	CreatedAt   time.Time      `gorm:"autoCreateTime"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
