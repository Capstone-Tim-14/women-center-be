package requests

type CounselingPackageRequest struct {
	Title            string `json:"title" validate:"required" form:"title"`
	Description      string `json:"description" validate:"required" form:"description"`
	Thumbnail        *string
	Session_per_week uint   `json:"session_per_week" validate:"required" form:"session_per_week"`
	Price            string `json:"price" validate:"required" form:"price"`
}
