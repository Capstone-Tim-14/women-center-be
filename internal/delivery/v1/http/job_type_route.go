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

func HttpJobTypeRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	JobTypeRepo := repositories.NewJobTypeRepository(db)
	JobTypeService := services.NewJobTypeService(JobTypeRepo, validate)
	JobTypeHandler := handlers.NewJobTypeHandler(JobTypeService)

	tag := group.Group("/admin/career/job-type", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	tag.POST("", JobTypeHandler.CreateJobType)
	tag.GET("", JobTypeHandler.ShowAllJobType)
}
