package schema

import (
	"time"

	"gorm.io/gorm"
)

type Admin struct {
	Id            uint `gorm:"primaryKey;"`
	Credential_id uint
	Credential    *Credentials   `gorm:"foreignKey:Credential_id;references:Id;"`
	First_name    string         `gorm:"type:varchar(100)"`
	Last_name     string         `gorm:"type:varchar(100)"`
	CreatedAt     time.Time      `gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime:mili"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
