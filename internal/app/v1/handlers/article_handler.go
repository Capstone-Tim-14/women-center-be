package handlers

import (
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	CreateArticle(ctx echo.Context) error
}

type ArticleHandlerImpl struct {
	ArticleService services.ArticleService
}

func NewArticleHandler(article services.ArticleService) ArticleHandler {
	return &ArticleHandlerImpl{
		ArticleService: article,
	}
}

func (handler *ArticleHandlerImpl) CreateArticle(ctx echo.Context) error {
	adminCreateRequest := requests.ArticleRequest{}
	err := ctx.Bind(&adminCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.ArticleService.CreateArticle(ctx, adminCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Article created successfully", nil)
}
