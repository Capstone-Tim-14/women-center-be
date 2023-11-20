package resources

type ArticleResource struct {
	Id          int    `json:"id,omitempty"`
	Title       string `json:"title,omitempty"`
	Thumbnail   string `json:"thumbnail,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Content     string `json:"content,omitempty"`
	Status      string `json:"status,omitempty"`
	Author      Author `json:"author,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
}

type Author struct {
	Name string `json:"name,omitempty"`
	Role string `json:"role,omitempty"`
}
