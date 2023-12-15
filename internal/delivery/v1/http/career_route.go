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
	JobTypeRepo := repositories.NewJobTypeRepository(db)
	CareerhasTypeRepo := repositories.NewCareerhasTypeRepository(db)
	CareerService := services.NewCareerService(services.CareerServiceImpl{
		CareerRepo:        CareerRepo,
		JobTypeRepo:       JobTypeRepo,
		CareerhasTypeRepo: CareerhasTypeRepo,
		Validator:         validate,
	})
	CareerHandler := handlers.NewCareerHandler(CareerService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))
	verifyTokenUser := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	CareerGroup := verifyTokenAdmin.Group("/career")

	CareerGroup.POST("", CareerHandler.CreateCareer)
	CareerGroup.GET("", CareerHandler.FindAllCareer)
	CareerGroup.GET("/:id", CareerHandler.FindDetailCareer)
	CareerGroup.POST("/:id/add-job-type", CareerHandler.AddJobType)
	CareerGroup.DELETE("/:id/remove-job-type", CareerHandler.RemoveJobType)
	CareerGroup.PUT("/:id", CareerHandler.UpdateCareer)
	CareerGroup.DELETE("/:id", CareerHandler.DeleteCareer)
	CareerGroup.GET("/recomendation", CareerHandler.RecomendationCareerList)
	CareerGroup.PUT("/recomended", CareerHandler.UpdateRecomendationCareer)

	verifyTokenUser.GET("/careers", CareerHandler.FindAllCareer)
	verifyTokenUser.GET("/career/:id", CareerHandler.FindDetailCareer)
	verifyTokenUser.GET("/recomendation-careers", CareerHandler.RecomendationCareerListForMobile)
}
