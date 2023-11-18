package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/requests/v1"
)

func TagArticleCreateRequestToTagArticleDomain(request requests.TagArticleRequest) *domain.Tag_Article {
	return &domain.Tag_Article{
		Name:        request.Name,
		Description: request.Description,
	}
}
