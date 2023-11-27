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
	CreatedAt             time.Time
	UpdatedAt             time.Time
}
