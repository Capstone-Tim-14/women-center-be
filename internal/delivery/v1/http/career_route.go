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

func HttpCareerRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	CareerRepo := repositories.NewCareerRepository(db)
	CareerService := services.NewCareerService(services.CareerServiceImpl{
		CareerRepo: CareerRepo,
		Validator:  validate,
	})
	CareerHandler := handlers.NewCareerHandler(CareerService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	CareerGroup := verifyTokenAdmin.Group("/career")

	CareerGroup.POST("", CareerHandler.CreateCareer)
	CareerGroup.GET("", CareerHandler.FindAllCareer)
	CareerGroup.GET("/:id", CareerHandler.FindDetailCareer)
	CareerGroup.PUT("/:id", CareerHandler.UpdateCareer)

}
