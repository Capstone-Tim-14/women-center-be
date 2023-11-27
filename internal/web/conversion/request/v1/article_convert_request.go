package conversion

import (
	"time"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"

	"github.com/gosimple/slug"
)

func ArticleCreateRequestToArticleDomain(request requests.ArticleRequest) *domain.Articles {
	return &domain.Articles{
		Title:         request.Title,
		Slug:          slug.Make(request.Title),
		Content:       request.Content,
		Thumbnail:     request.Thumbnail,
		PublishedAt:   time.Now(),
		Admin_id:      request.Admin_id,
		Counselors_id: request.Counselors_id,
		Status:        "REVIEW",
	}
}
func ArticleUpdateRequestToArticleDomain(request requests.ArticleRequest) *domain.Articles {
	article := &domain.Articles{
		Title:   request.Title,
		Slug:    slug.Make(request.Title),
		Content: request.Content,
	}
	if request.Thumbnail != nil {
		article.Thumbnail = request.Thumbnail
	}
	return article
}
