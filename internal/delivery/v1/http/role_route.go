package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpRoleRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	RoleService := services.NewRoleService(RoleRepo, validate)
	RoleHandler := handlers.NewRoleHandler(RoleService)

	role := group.Group("/roles")

	role.POST("/create", RoleHandler.CreateRoleHandler)

}
