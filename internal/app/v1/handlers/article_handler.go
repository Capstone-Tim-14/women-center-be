package handlers

import (
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type ArticleHandler interface {
	CreateArticle(ctx echo.Context) error
	FindAllArticle(ctx echo.Context) error
	FindArticleBySlug(ctx echo.Context) error
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

func (handler *ArticleHandlerImpl) FindAllArticle(ctx echo.Context) error {
	response, err := handler.ArticleService.FindAllArticle(ctx)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	articleResponse := conversion.ConvertArticleResource(response)

	return responses.StatusOK(ctx, "Success Get All Article", articleResponse)
}

func (handler *ArticleHandlerImpl) FindArticleBySlug(ctx echo.Context) error {
	slug := ctx.Param("slug")
	response, err := handler.ArticleService.FindArticleBySlug(ctx, slug)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	articleResponse := conversion.ConvertSingleArticleResource(response)

	return responses.StatusOK(ctx, "Success Get Article", articleResponse)
}
