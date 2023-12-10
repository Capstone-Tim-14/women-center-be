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

func HttpUserRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	UserRepo := repositories.NewUserRepository(db)
	ArticleRepo := repositories.NewArticleRepository(db)
	FavoriteRepo := repositories.NewArticleFavoriteRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	CounselorFavoriteRepo := repositories.NewCounselorFavoriteRepository(db)
	UserService := services.NewUserService(services.UserServiceImpl{
		RoleRepo:              RoleRepo,
		UserRepo:              UserRepo,
		CounselorRepo:         CounselorRepo,
		CounselorFavoriteRepo: CounselorFavoriteRepo,
		Validator:             validate,
		ArticleRepo:           ArticleRepo,
		FavoriteArticle:       FavoriteRepo,
	})
	UserHandler := handlers.NewUserHandler(UserService)

	user := group.Group("/users")

	user.POST("/register", UserHandler.RegisterHandler)

	userVerify := user.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	userVerify.GET("/profile", UserHandler.ProfileHandler)
	userVerify.PUT("/profile", UserHandler.UpdateProfileHandler)

	verifyTokenFavorite := group.Group("/article", middlewares.VerifyTokenSignature("SECRET_KEY"))
	verifyTokenFavorite.POST("/:slug/add-favorite", UserHandler.AddFavoriteArticleHandler)
	verifyTokenFavorite.DELETE("/:slug/delete-favorite", UserHandler.DeleteFavoriteArticleHandler)
	verifyTokenFavorite.GET("/favorites", UserHandler.AllFavoriteArticleHandler)

	verifyCounselorFavorite := group.Group("/counselor", middlewares.VerifyTokenSignature("SECRET_KEY"))
	verifyCounselorFavorite.POST("/:id/add-counselor-favorite", UserHandler.AddCounselorFavorite)
	verifyCounselorFavorite.DELETE("/:id/remove-counselor-favorite", UserHandler.RemoveCounselorFavorite)
	verifyCounselorFavorite.GET("/favorites", UserHandler.GetCounselorFavorite)
}
