package routes

import (
	"woman-center-be/internal/app/handlers"
	"woman-center-be/internal/app/repositories"
	"woman-center-be/internal/app/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpUserRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	UserRepo := repositories.NewUserRepository(db)
	UserService := services.NewUserService(UserRepo, validate)
	UserHandler := handlers.NewUserHandler(UserService)

	user := group.Group("/users")

	user.POST("/register", UserHandler.RegisterHandler)

}
