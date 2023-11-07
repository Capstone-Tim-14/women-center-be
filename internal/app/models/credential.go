package models

import (
	"time"

	"gorm.io/gorm"
)

type Credentials struct {
	Id        uint           `gorm:"primaryKey;"`
	Username  string         `gorm:"type:varchar(100)"`
	Password  string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
