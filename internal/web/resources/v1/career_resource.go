package resources

type CareerResource struct {
	Id                    uint      `json:"id,omitempty"`
	Title_job             string    `json:"title_job,omitempty"`
	Logo                  string    `json:"logo,omitempty"`
	Company_name          string    `json:"company_name,omitempty"`
	PublishedAt           string    `json:"published_at,omitempty"`
	Cover                 string    `json:"cover,omitempty"`
	Linkedin_url          string    `json:"linkedin_url,omitempty"`
	Location              string    `json:"location,omitempty"`
	About_job             string    `json:"about_job,omitempty"`
	About_company         string    `json:"about_company,omitempty"`
	Required_skill        string    `json:"required_skill,omitempty"`
	Company_industry      string    `json:"company_industry,omitempty"`
	Size_company_employee uint      `json:"size_company_employee,omitempty"`
	Job_type              []JobType `json:"job_type,omitempty"`
	Recomendation         string    `json:"recomendation,omitempty"`
	CreatedAt             string    `json:"created_at,omitempty"`
	UpdatedAt             string    `json:"updated_at,omitempty"`
}

type JobType struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}
