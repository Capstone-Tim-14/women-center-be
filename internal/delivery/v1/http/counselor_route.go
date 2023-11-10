package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpCounselorRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	CounselorService := services.NewCounselorService(CounselorRepo, validate, RoleRepo)
	CounselorHandler := handlers.NewCounselorHandler(CounselorService)

	counselor := group.Group("/counselors")

	counselor.POST("/register", CounselorHandler.RegisterHandler)

}
