package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpAdminRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	AdminRepo := repositories.NewAdminRepository(db)
	AdminService := services.NewAdminService(AdminRepo, validate, RoleRepo)
	AdminHandler := handlers.NewAdminHandler(AdminService)

	admin := group.Group("/admin")

	admin.POST("/register", AdminHandler.RegisterHandler)

}
