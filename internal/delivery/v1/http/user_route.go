package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpUserRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	UserRepo := repositories.NewUserRepository(db)
	UserService := services.NewUserService(UserRepo, RoleRepo, validate)
	UserHandler := handlers.NewUserHandler(UserService)

	user := group.Group("/users")

	user.POST("/register", UserHandler.RegisterHandler)

}
