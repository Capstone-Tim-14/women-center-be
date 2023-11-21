package services

import (
	"fmt"
	"mime/multipart"
	"strconv"
	"woman-center-be/internal/app/v1/models/domain"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/pkg/query"
	"woman-center-be/pkg/storage"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ArticleService interface {
	CreateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) (*domain.Articles, []exceptions.ValidationMessage, error)
	FindAllArticle(ctx echo.Context) ([]domain.Articles, *query.Pagination, error)
	DeleteArticle(ctx echo.Context) error
}

type ArticleServiceImpl struct {
	ArticleRepo   repositories.ArticleRepository
	AdminRepo     repositories.AdminRepository
	CounselorRepo repositories.CounselorRepository
	validator     *validator.Validate
}

func NewArticleService(article repositories.ArticleRepository, validator *validator.Validate, admin repositories.AdminRepository, counselor repositories.CounselorRepository) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepo:   article,
		AdminRepo:     admin,
		CounselorRepo: counselor,
		validator:     validator,
	}
}

func (service *ArticleServiceImpl) CreateArticle(ctx echo.Context, request requests.ArticleRequest, thumbnail *multipart.FileHeader) (*domain.Articles, []exceptions.ValidationMessage, error) {
	err := service.validator.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err), nil
	}

	author := helpers.GetAuthClaims(ctx)

	if author.Role == "admin" || author.Role == "super_admin" {
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

	ThumbnailCloudURL, errUploadThumbnail := storage.DropboxUploadEndpoint(thumbnail, "articles")

	if errUploadThumbnail != nil {
		return nil, nil, errUploadThumbnail
	}

	request.Thumbnail = ThumbnailCloudURL

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

	Page, _ := strconv.Atoi(QueryPage)
	Limit, _ := strconv.Atoi(QueryLimit)

	Paginate := query.Pagination{
		Page:  uint(Page),
		Limit: uint(Limit),
	}

	result, paginate, err := service.ArticleRepo.FindAllArticle(orderBy, Paginate)
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
