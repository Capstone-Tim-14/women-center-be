package domain

import "time"

type Career struct {
	Id                    uint
	Title_job             string
	Logo                  string
	Company_name          string
	PublishedAt           time.Time
	Cover                 string
	Linkedin_url          string
	Location              string
	About_job             string
	About_company         string
	Required_skill        string
	Company_industry      string
	Size_company_employee uint
	Job_type              []Job_Type `gorm:"many2many:career_has_types;foreignKey:Id;references:Id;"`
	Recomendation         bool       `gorm:"default:false"`
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
