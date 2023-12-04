package resources

type CounselingPackageResource struct {
	Id                 uint   `json:"id,omitempty"`
	Package_name       string `json:"package_name,omitempty"`
	Description        string `json:"description,omitempty"`
	Thumbnail          string `json:"thumbnail,omitempty"`
	Number_of_sessions uint   `json:"number_of_sessions,omitempty"`
	Price              string `json:"price,omitempty"`
	PublishedAt        string `json:"published_at,omitempty"`
	CreatedAt          string `json:"created_at,omitempty"`
	UpdatedAt          string `json:"updated_at,omitempty"`
	DeletedAt          string `json:"deleted_at,omitempty"`
}
