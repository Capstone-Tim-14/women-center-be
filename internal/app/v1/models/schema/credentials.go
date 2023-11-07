package schema

import (
	"time"

	"gorm.io/gorm"
)

type Credentials struct {
	Id        uint `gorm:"primaryKey;"`
	Role_id   uint
	Role      *Roles         `gorm:"foreignKey:Role_id;references:Id"`
	Username  string         `gorm:"type:varchar(100)"`
	Password  string         `gorm:"type:varchar(255)"`
	CreatedAt time.Time      `gorm:"autoCreateTime"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
