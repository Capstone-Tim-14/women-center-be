package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpAuthRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	UserRepo := repositories.NewUserRepository(db)

	AuthService := services.NewAuthService(RoleRepo, UserRepo, CounselorRepo, validate)
	AuthHandler := handlers.NewAuthHandler(AuthService)

	group.GET("/google-auth", AuthHandler.OauthGoogleHandler)
	group.GET("/callback-google-auth", AuthHandler.OauthCallbackGoogleHandler)

	Users := group.Group("/user")
	Counselor := group.Group("/counselor")

	Users.POST("/auth", AuthHandler.UserAuthHandler)
	Counselor.POST("/auth", AuthHandler.CounselorAuthHandler)

}
