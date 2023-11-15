package services

import (
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/web/requests"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleService interface {
	CreateArticle(ctx echo.Context, request requests.ArticleRequest) (*domain.Articles, []exceptions.ValidationMessage, error)
}

type ArticleServiceImpl struct {
	ArticleRepo repositories.ArticleRepository
	validator   *validator.Validate
}

func NewArticleService(article repositories.ArticleRepository, validator *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepo: article,
		validator:   validator,
	}
}

func (service *ArticleServiceImpl) CreateArticle(ctx echo.Context, request requests.ArticleRequest) (*domain.Articles, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

}
