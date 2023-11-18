package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type TagService interface {
	CreateTag(ctx echo.Context, request requests.TagArticleRequest) (*domain.Tag_Article, []exceptions.ValidationMessage, error)
}

type TagServiceImpl struct {
	TagRepo   repositories.TagRepository
	validator *validator.Validate
}

func NewTagService(tagRepo repositories.TagRepository, validator *validator.Validate) TagService {
	return &TagServiceImpl{
		TagRepo:   tagRepo,
		validator: validator,
	}
}

func (service *TagServiceImpl) CreateTag(ctx echo.Context, request requests.TagArticleRequest) (*domain.Tag_Article, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	// Check if the tag already exists
	existingTag, _ := service.TagRepo.FindTagByName(request.Name)
	if existingTag != nil {
		return nil, nil, fmt.Errorf("tag already exists")
	}

	newTag := conversion.TagArticleCreateRequestToTagArticleDomain(request)

	// Create the new tag
	createdTag, err := service.TagRepo.CreateTag(newTag)
	if err != nil {
		return nil, nil, fmt.Errorf("error while creating tag: %s", err.Error())
	}

	return createdTag, nil, nil
}
