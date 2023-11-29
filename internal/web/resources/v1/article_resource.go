package resources

type ArticleResource struct {
	Id          uint              `json:"id,omitempty"`
	Title       string            `json:"title,omitempty"`
	Thumbnail   string            `json:"thumbnail,omitempty"`
	Slug        string            `json:"slug,omitempty"`
	Content     string            `json:"content,omitempty"`
	Status      string            `json:"status,omitempty"`
	Author      Author            `json:"author,omitempty"`
	Tag         []ArticleCategory `json:"category,omitempty"`
	CreatedAt   string            `json:"created_at,omitempty"`
	UpdatedAt   string            `json:"updated_at,omitempty"`
	PublishedAt string            `json:"published_at,omitempty"`
	TimeUpload  string            `json:"time_upload,omitempty"`
}

type ArticleCounseloResource struct {
	Article_publish  int               `json:"article_publish"`
	Article_review   int               `json:"Article_review"`
	Article_rejected int               `json:"Article_rejected"`
	ArticleList      []ArticleResource `json:"lists,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}

type ArticleCategory struct {
	Id   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
