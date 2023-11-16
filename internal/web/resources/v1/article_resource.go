package resources

type ArticleResource struct {
	Id         int    `json:"id,omitempty"`
	Title      string `json:"title,omitempty"`
	Thumbnail  string `json:"thumbnail,omitempty"`
	Slug       string `json:"slug,omitempty"`
	Content    string `json:"content,omitempty"`
	Status     string `json:"status,omitempty"`
	Admin      *AdminResource
	Counselors *CounselorResource
}
