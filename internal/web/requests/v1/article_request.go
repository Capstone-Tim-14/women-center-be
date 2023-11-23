package requests

type ArticleRequest struct {
	Title         string `json:"title" validate:"required" form:"title"`
	Content       string `json:"content" validate:"required" form:"content"`
	Thumbnail     *string
	Admin_id      *uint
	Counselors_id *uint
}
