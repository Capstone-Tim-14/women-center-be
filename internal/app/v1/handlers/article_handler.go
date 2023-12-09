package handlers

import (
	"fmt"
	"strconv"
	"strings"
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
	DeleteArticle(ctx echo.Context) error
	UpdatePublishedArticle(ctx echo.Context) error
	UpdateArticle(ctx echo.Context) error
	AddTagArticle(ctx echo.Context) error
	RemoveTagArticle(ctx echo.Context) error
	LatestArticleHandler(ctx echo.Context) error
	FindAllArticleUser(ctx echo.Context) error
	FindArticleBySlugForUser(ctx echo.Context) error
	FindArticleBySlug(ctx echo.Context) error
	AllArticleCounselorHandler(ctx echo.Context) error
}

type ArticleHandlerImpl struct {
	ArticleService services.ArticleService
}

func NewArticleHandler(article services.ArticleService) ArticleHandler {
	return &ArticleHandlerImpl{
		ArticleService: article,
	}
}

func (handler *ArticleHandlerImpl) FindArticleBySlugForUser(ctx echo.Context) error {

	slug := ctx.Param("slug")
	response, err := handler.ArticleService.FindArticleForUserBySlug(ctx, slug)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	articleResponse := conversion.ConvertSingleArticleResource(response)

	return responses.StatusOK(ctx, "Success Get Article", articleResponse)

}

func (handler *ArticleHandlerImpl) AllArticleCounselorHandler(ctx echo.Context) error {

	response, err := handler.ArticleService.FindAllArticleCounselor(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Articles empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success Get Article Counselor", response)
}

func (handler *ArticleHandlerImpl) RemoveTagArticle(ctx echo.Context) error {

	var request requests.ArticleHasManyRequest

	GetId := ctx.Param("id")

	ErrBinding := ctx.Bind(&request)

	if ErrBinding != nil {
		return exceptions.StatusBadRequest(ctx, ErrBinding)
	}

	ParseToId, errParsing := strconv.Atoi(GetId)

	if errParsing != nil {
		fmt.Errorf(errParsing.Error())
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Invalid Format id"))
	}

	Validation, ErrGetArticle := handler.ArticleService.RemoveTagArticle(ctx, ParseToId, request)

	if Validation != nil {
		return exceptions.ValidationException(ctx, "Validation Error", Validation)
	}

	if ErrGetArticle != nil {
		if strings.Contains(ErrGetArticle.Error(), "Article not found") {
			return exceptions.StatusNotFound(ctx, ErrGetArticle)
		}
		if strings.Contains(ErrGetArticle.Error(), "One of article request is not found") {
			return exceptions.StatusNotFound(ctx, ErrGetArticle)
		}

		return exceptions.StatusInternalServerError(ctx, ErrGetArticle)
	}

	return responses.StatusOK(ctx, "Article Remove Successfully", nil)

}

func (handler *ArticleHandlerImpl) LatestArticleHandler(ctx echo.Context) error {

	GetLatestResponse, err := handler.ArticleService.GetLatestArticle()

	if err != nil {
		if strings.Contains(err.Error(), "Article not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success get latest article", GetLatestResponse)

}

func (handler *ArticleHandlerImpl) CreateArticle(ctx echo.Context) error {
	adminCreateRequest := requests.ArticleRequest{}
	Thumbnail, errThumb := ctx.FormFile("thumbnail")

	if errThumb != nil {
		return exceptions.StatusBadRequest(ctx, errThumb)
	}

	err := ctx.Bind(&adminCreateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}
	_, validation, err := handler.ArticleService.CreateArticle(ctx, adminCreateRequest, Thumbnail)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Title already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Article created successfully", nil)
}

func (handler *ArticleHandlerImpl) FindAllArticleUser(ctx echo.Context) error {

	GetArticleResource, err := handler.ArticleService.FindAllArticleUser()

	if err != nil {
		if strings.Contains(err.Error(), "Article is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success Get All Article", GetArticleResource)
}

func (handler *ArticleHandlerImpl) FindAllArticle(ctx echo.Context) error {

	response, meta, err := handler.ArticleService.FindAllArticle(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "Articles is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	articleResponse := conversion.ConvertArticleResource(response)

	return responses.StatusOKWithMeta(ctx, "Success Get All Article", meta, articleResponse)
}

func (handler *ArticleHandlerImpl) DeleteArticle(ctx echo.Context) error {

	err := handler.ArticleService.DeleteArticle(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "Article not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Article deleted successfully", nil)
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

func (handler *ArticleHandlerImpl) UpdatePublishedArticle(ctx echo.Context) error {
	request := requests.PublishArticle{}
	err := ctx.Bind(&request)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	validation, err := handler.ArticleService.UpdatePublishedArticle(ctx, request)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Article not found") {
			return exceptions.StatusNotFound(ctx, err)
		}

		if strings.Contains(err.Error(), "Article already published") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		if strings.Contains(err.Error(), "Article already rejected") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	if request.Status == "APPROVED" {
		return responses.StatusCreated(ctx, "Congratulations, the article is PUBLISH", nil)
	} else if request.Status == "REJECTED" {
		return responses.StatusCreated(ctx, "Sorry, the Article is REJECTED", nil)
	}

	return responses.StatusCreated(ctx, "The article status updated", nil)
}

func (handler *ArticleHandlerImpl) AddTagArticle(ctx echo.Context) error {

	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)
	var request requests.ArticlehasTagRequest
	errBinding := ctx.Bind(&request)

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.ArticleService.AddTagArticle(ctx, convertid, request)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success add category to article", nil)
}

func (handler *ArticleHandlerImpl) UpdateArticle(ctx echo.Context) error {

	var request requests.ArticleRequest
	errBinding := ctx.Bind(&request)

	Thumbnail, _ := ctx.FormFile("thumbnail")

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.ArticleService.UpdateArticle(ctx, request, Thumbnail)
	if validation != nil {
		return exceptions.ValidationException(ctx, "Error Validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Article updated!", nil)
}
