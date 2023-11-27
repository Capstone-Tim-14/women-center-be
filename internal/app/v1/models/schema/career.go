package schema

import "time"

type Career struct {
	Id                    uint      `gorm:"primaryKey;"`
	Title_job             string    `gorm:"type:varchar(100)"`
	Logo                  string    `gorm:"type:varchar(255)"`
	Company_name          string    `gorm:"type:varchar(255)"`
	PublishedAt           time.Time `gorm:"autoCreateTime"`
	Cover                 string    `gorm:"type:varchar(255)"`
	Linkedin_url          string    `gorm:"type:varchar(255)"`
	Location              string    `gorm:"type:varchar(255)"`
	About_job             string
	About_company         string
	Required_skill        string `gorm:"type:varchar(255)"`
	Company_industry      string `gorm:"type:varchar(255)"`
	Size_company_employee uint
	CreatedAt             time.Time `gorm:"autoCreateTime"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime"`
}
