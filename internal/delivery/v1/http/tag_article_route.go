package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/delivery/v1/middlewares"

	_ "github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpTagRoute(group *echo.Group, db *gorm.DB) {

	TagRepo := repositories.NewTagRepository(db)
	TagService := services.NewTagService(TagRepo)
	TagHandler := handlers.NewTagHandler(TagService)

	tag := group.Group("/admin/articles/category", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	tag.POST("", TagHandler.CreateTag)
	// tag.GET("", TagHandler.FindtagsHandler)
	// tag.DELETE("/delete/:id", tagHandler.DeletetagByIdHandler)

}
