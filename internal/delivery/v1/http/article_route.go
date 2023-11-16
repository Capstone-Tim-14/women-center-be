package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/delivery/v1/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpArticleRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	ArticleRepo := repositories.NewArticleRepository(db)
	AdminRepo := repositories.NewAdminRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	ArticleService := services.NewArticleService(ArticleRepo, validate, AdminRepo, CounselorRepo)
	ArticleHandler := handlers.NewArticleHandler(ArticleService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))
	varifyTokenCounselor := group.Group("/counselor", middlewares.VerifyTokenSignature("SECRET_KEY"))

	articleAdmin := verifyTokenAdmin.Group("/articles")
	articleAdmin.POST("", ArticleHandler.CreateArticle)
	articleAdmin.GET("", ArticleHandler.FindAllArticle)
	articleAdmin.GET("/:slug", ArticleHandler.FindArticleBySlug)

	articleCounselor := varifyTokenCounselor.Group("/articles")
	articleCounselor.POST("", ArticleHandler.CreateArticle)
}
