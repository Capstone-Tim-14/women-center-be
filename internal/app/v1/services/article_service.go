package services

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	conResources "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/pkg/query"
	"woman-center-be/pkg/storage"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleService interface {
	CreateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) (*domain.Articles, []exceptions.ValidationMessage, error)
	FindAllArticleUser() ([]resources.ArticleResource, error)
	FindAllArticle(ctx echo.Context) ([]domain.Articles, *query.Pagination, error)
	DeleteArticle(ctx echo.Context) error
	UpdatePublishedArticle(ctx echo.Context, request requests.PublishArticle) ([]exceptions.ValidationMessage, error)
	FindArticleBySlug(ctx echo.Context, slug string) (*domain.Articles, error)
	AddTagArticle(ctx echo.Context, id int, request requests.ArticlehasTagRequest) ([]exceptions.ValidationMessage, error)
	UpdateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) ([]exceptions.ValidationMessage, error)
}

type ArticleServiceImpl struct {
	ArticleRepo       repositories.ArticleRepository
	AdminRepo         repositories.AdminRepository
	CounselorRepo     repositories.CounselorRepository
	Validator         *validator.Validate
	TagRepo           repositories.TagRepository
	ArticlehasTagRepo repositories.ArticlehasTagRepository
}

func NewArticleService(articleServiceImpl ArticleServiceImpl) ArticleService {
	return &articleServiceImpl
}

func (service *ArticleServiceImpl) FindAllArticleUser() ([]resources.ArticleResource, error) {

	Getarticles, err := service.ArticleRepo.GetListArticleForUser()

	if err != nil {
		return nil, err
	}

	Resources := conResources.ConvertArticleResource(Getarticles)

	return Resources, nil

}

func (service *ArticleServiceImpl) CreateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) (*domain.Articles, []exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	existingTitle, _ := service.ArticleRepo.FindByTitle(request.Title)
	if existingTitle != nil {
		return nil, nil, fmt.Errorf("Title already exists")
	}

	author := helpers.GetAuthClaims(ctx)

	if author.Role == "admin" || author.Role == "super admin" {
		admin, err := service.AdminRepo.FindyByEmail(author.Email)
		if err != nil {
			return nil, nil, fmt.Errorf("Error get admin")
		}

		request.Admin_id = &admin.Id
	} else if author.Role == "counselor" {
		counselor, err := service.CounselorRepo.FindyByEmail(author.Email)
		if err != nil {
			return nil, nil, fmt.Errorf("Error get counselor")
		}

		request.Counselors_id = &counselor.Id
	}

	ThumbnailCloudURL, errUploadThumbnail := storage.S3PutFile(thumbnail, "articles/thumbnail")

	if errUploadThumbnail != nil {
		return nil, nil, errUploadThumbnail
	}

	request.Thumbnail = &ThumbnailCloudURL

	article := conversion.ArticleCreateRequestToArticleDomain(request)

	result, err := service.ArticleRepo.CreateArticle(article)
	if err != nil {
		return nil, nil, fmt.Errorf("Error create article")
	}

	return result, nil, nil
}

func (service *ArticleServiceImpl) FindAllArticle(ctx echo.Context) ([]domain.Articles, *query.Pagination, error) {

	orderBy := ctx.QueryParam("orderBy")
	QueryLimit := ctx.QueryParam("limit")
	QueryPage := ctx.QueryParam("page")
	Search := ctx.QueryParam("search")

	Page, _ := strconv.Atoi(QueryPage)
	Limit, _ := strconv.Atoi(QueryLimit)

	Paginate := query.Pagination{
		Page:  uint(Page),
		Limit: uint(Limit),
	}

	result, paginate, err := service.ArticleRepo.FindAllArticle(orderBy, Search, Paginate)
	if err != nil {
		return nil, nil, fmt.Errorf("Article is empty")
	}

	return result, paginate, nil
}

func (service *ArticleServiceImpl) DeleteArticle(ctx echo.Context) error {
	id := ctx.Param("id")
	getId, _ := strconv.Atoi(id)
	err := service.ArticleRepo.DeleteArticleById(getId)
	if err != nil {
		return err
	}

	return nil
}

func (service *ArticleServiceImpl) UpdatePublishedArticle(ctx echo.Context, request requests.PublishArticle) ([]exceptions.ValidationMessage, error) {

	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	slug := ctx.Param("slug")

	findSlug, err := service.ArticleRepo.FindBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("Article not found")

	}
	if findSlug.Status == "Published" {
		return nil, fmt.Errorf("Article already published")
	}

	if findSlug.Status == "Rejected" {
		return nil, fmt.Errorf("Article already rejected")
	}

	var errUpdateStatus error

	if request.Status == "APPROVED" {
		errUpdateStatus = service.ArticleRepo.UpdateStatusArticle(slug, "PUBLISHED")
	} else {
		errUpdateStatus = service.ArticleRepo.UpdateStatusArticle(slug, "REJECTED")
	}

	if errUpdateStatus != nil {
		return nil, fmt.Errorf("Error update status article")
	}

	return nil, nil
}

func (service *ArticleServiceImpl) AddTagArticle(ctx echo.Context, id int, request requests.ArticlehasTagRequest) ([]exceptions.ValidationMessage, error) {

	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	article, errArticle := service.ArticleRepo.FindById(id)
	if errArticle != nil {
		return nil, errArticle
	}

	tag, errTag := service.TagRepo.FindTagByName(request.Name)
	if errTag != nil {
		return nil, errTag
	}

	errAddTag := service.ArticlehasTagRepo.AddTag(*article, tag)
	if errAddTag != nil {
		return nil, errAddTag
	}

	return nil, nil
}

func (service *ArticleServiceImpl) FindArticleBySlug(ctx echo.Context, slug string) (*domain.Articles, error) {
	result, err := service.ArticleRepo.FindBySlug(slug)
	if err != nil {
		return nil, fmt.Errorf("Article not found")
	}

	return result, nil
}

func (service *ArticleServiceImpl) UpdateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) ([]exceptions.ValidationMessage, error) {
	err := service.Validator.Struct(request)
	if err != nil {
		return helpers.ValidationError(ctx, err), nil
	}

	id := ctx.Param("id")
	getId, _ := strconv.Atoi(id)

	_, err = service.ArticleRepo.FindById(getId)
	if err != nil {
		return nil, fmt.Errorf("Article not found")

	}
	if thumbnail != nil {

		ThumbnailCloudURL, errUploadThumbnail := storage.S3PutFile(thumbnail, "articles/thumbnail")

		if errUploadThumbnail != nil {
			return nil, errUploadThumbnail
		}

		request.Thumbnail = &ThumbnailCloudURL
	}
	article := conversion.ArticleUpdateRequestToArticleDomain(request)

	_, err = service.ArticleRepo.UpdateArticle(getId, article), nil
	if err != nil {
		return nil, fmt.Errorf("Error update article")
	}

	return nil, nil
}
