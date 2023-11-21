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

func HttpSpecialistRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	SpecialistRepo := repositories.NewSpecialistRepository(db)
	SpecialistService := services.NewSpecialistService(SpecialistRepo, validate)
	SpecialistHandler := handlers.NewSpecialistHandler(SpecialistService)

	specialist := group.Group("/admin/specialist", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	specialist.POST("", SpecialistHandler.CreateSpecialist)
	//specialist.GET("", TagHandler.GetAllTagsHandler)
	//specialist.DELETE("/:id", TagHandler.DeleteTagByIdHandler)

}
