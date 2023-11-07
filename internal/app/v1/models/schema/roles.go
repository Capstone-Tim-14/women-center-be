package schema

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	Id        uint           `gorm:"primaryKey;"`
	Name      string         `gorm:"type:varchar(50)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:milli"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
