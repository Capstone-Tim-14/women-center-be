package requests

type ArticleRequest struct {
	Title         string `json:"title" validate:"required" form:"title"`
	Content       string `json:"content" validate:"required" form:"content"`
	Thumbnail     *string
	Admin_id      *uint
	Counselors_id *uint
}

type PublishArticle struct {
	Status string `json:"status" validate:"required" form:"status"`
}
