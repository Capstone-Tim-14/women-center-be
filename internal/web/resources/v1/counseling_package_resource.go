package resources

type CounselingPackageResource struct {
	Id               uint   `json:"id,omitempty"`
	Title            string `json:"title,omitempty"`
	Thumbnail        string `json:"thumbnail,omitempty"`
	Session_per_week uint   `json:"session_per_week,omitempty"`
	Price            string `json:"price,omitempty"`
	Description      string `json:"description,omitempty"`
	CreatedAt        string `json:"created_at,omitempty"`
	UpdatedAt        string `json:"updated_at,omitempty"`
	DeletedAt        string `json:"deleted_at,omitempty"`
}
