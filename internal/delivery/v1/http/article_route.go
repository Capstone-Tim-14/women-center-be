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
	TagRepo := repositories.NewTagRepository(db)
	ArticlehasTagRepo := repositories.NewArticlehasTagRepository(db)
	ArticleService := services.NewArticleService(services.ArticleServiceImpl{
		ArticleRepo:       ArticleRepo,
		AdminRepo:         AdminRepo,
		CounselorRepo:     CounselorRepo,
		TagRepo:           TagRepo,
		ArticlehasTagRepo: ArticlehasTagRepo,
		Validator:         validate,
	})
	ArticleHandler := handlers.NewArticleHandler(ArticleService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))
	verifyToken := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	CounselorGroup := verifyToken.Group("/counselor")

	articleAdmin := verifyTokenAdmin.Group("/articles")
	articleAdmin.POST("", ArticleHandler.CreateArticle)

	articleAdmin.GET("", ArticleHandler.FindAllArticle).Name = "admin.articles.get-all"
	articleAdmin.DELETE("/:id", ArticleHandler.DeleteArticle)
	articleAdmin.GET("/:slug", ArticleHandler.FindArticleBySlug)
	articleAdmin.PUT("/approved-request/:slug", ArticleHandler.UpdatePublishedArticle)
	articleAdmin.POST("/:id/add-category", ArticleHandler.AddTagArticle)
	articleAdmin.DELETE("/:id/remove-category", ArticleHandler.RemoveTagArticle)
	articleAdmin.PUT("/:id", ArticleHandler.UpdateArticle)

	articleCounselor := CounselorGroup.Group("/articles")
	articleCounselor.POST("", ArticleHandler.CreateArticle)
	articleCounselor.GET("", ArticleHandler.AllArticleCounselorHandler)
	articleCounselor.PUT("/:id", ArticleHandler.UpdateArticle)
	articleCounselor.POST("/:id/add-category", ArticleHandler.AddTagArticleCounselor)
	articleCounselor.DELETE("/:id/remove-category", ArticleHandler.RemoveTagArticleCounselor)
	articleCounselor.DELETE("/:id", ArticleHandler.DeleteArticle)

	articleUser := verifyToken.Group("/articles")
	articleUser.GET("", ArticleHandler.FindAllArticleUser)
	verifyToken.GET("/article/latest", ArticleHandler.LatestArticleHandler)
	verifyToken.GET("/article/:slug", ArticleHandler.FindArticleBySlugForUser)
}
