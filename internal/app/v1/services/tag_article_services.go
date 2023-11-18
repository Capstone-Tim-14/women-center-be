package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"

	"github.com/labstack/echo/v4"
)

type TagService interface {
	CreateTag(ctx echo.Context, request requests.TagArticleRequest) (*domain.Tag_Article, error)
}

type TagServiceImpl struct {
	TagRepo repositories.TagRepository
}

func NewTagService(tagRepo repositories.TagRepository) TagService {
	return &TagServiceImpl{
		TagRepo: tagRepo,
	}
}

func (service *TagServiceImpl) CreateTag(ctx echo.Context, request requests.TagArticleRequest) (*domain.Tag_Article, error) {
	// Check if the tag already exists
	existingTag, _ := service.TagRepo.FindTagByName(request.Name)
	if existingTag != nil {
		return nil, fmt.Errorf("tag already exists")
	}

	newTag := conversion.TagArticleCreateRequestToTagArticleDomain(request)

	// Create the new tag
	createdTag, err := service.TagRepo.CreateTag(newTag)
	if err != nil {
		return nil, fmt.Errorf("error while creating tag: %s", err.Error())
	}

	return createdTag, nil
}
