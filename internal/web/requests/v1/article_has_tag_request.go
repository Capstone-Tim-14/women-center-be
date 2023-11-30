package requests

type ArticlehasTagRequest struct {
	Name string `json:"name" validate:"required"`
}

type ArticleHasManyRequest struct {
	Name []string `json:"name" validate:"required"`
}
