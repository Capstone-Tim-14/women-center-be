package requests

type CareerRequest struct {
	Title_job             string  `json:"title_job" validate:"required" form:"title_job"`
	Logo                  *string `json:"logo" validate:"required" form:"logo"`
	Company_name          string  `json:"company_name" validate:"required" form:"company_name"`
	Cover                 *string `json:"cover" validate:"required" form:"cover"`
	Linkedin_url          string  `json:"linkedin_url" validate:"required" form:"linkedin_url"`
	Location              string  `json:"location" validate:"required" form:"location"`
	About_job             string  `json:"about_job" validate:"required" form:"about_job"`
	About_company         string  `json:"about_company" validate:"required" form:"about_company"`
	Required_skill        string  `json:"required_skill" validate:"required" form:"required_skill"`
	Company_industry      string  `json:"company_industry" validate:"required" form:"company_industry"`
	Size_company_employee uint    `json:"size_company_employee" validate:"required" form:"size_company_employee"`
}

type CareerFilterRequest struct {
	JobType []string
}
