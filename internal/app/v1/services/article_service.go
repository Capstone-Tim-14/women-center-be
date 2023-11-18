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

type ArticleService interface {
	CreateArticle(ctx echo.Context, request requests.ArticleRequest) (*domain.Articles, []exceptions.ValidationMessage, error)
	FindAllArticle(ctx echo.Context) ([]domain.Articles, error)
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

func (service *ArticleServiceImpl) CreateArticle(ctx echo.Context, request requests.ArticleRequest) (*domain.Articles, []exceptions.ValidationMessage, error) {
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

	article := conversion.ArticleCreateRequestToArticleDomain(request)

	result, err := service.ArticleRepo.CreateArticle(article)
	if err != nil {
		return nil, nil, fmt.Errorf("Error create article")
	}

	return result, nil, nil
}

func (service *ArticleServiceImpl) FindAllArticle(ctx echo.Context) ([]domain.Articles, error) {

	orderBy := ctx.QueryParam("orderBy")

	result, err := service.ArticleRepo.FindAllArticle(orderBy)
	if err != nil {
		return nil, fmt.Errorf("Article not found")
	}

	return result, nil
}
