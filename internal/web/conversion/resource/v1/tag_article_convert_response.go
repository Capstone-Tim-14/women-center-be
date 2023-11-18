package conversion

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/web/resources/v1"
)

func TagArticleDomainToTagArticleResponse(tag *domain.Tag_Article) resources.TagArticleResource {
	return resources.TagArticleResource{
		Id:          tag.Id,
		Name:        tag.Name,
		Description: tag.Description,
	}
}

func ConvertTagResource(tags []domain.Tag_Article) []resources.TagArticleResource {
	tagResource := []resources.TagArticleResource{}
	for _, tag := range tags {
		tagResource = append(tagResource, resources.TagArticleResource{
			Id:          tag.Id,
			Name:        tag.Name,
			Description: tag.Description,
		})
	}

	return tagResource
}
