package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/delivery/v1/middlewares"
	"woman-center-be/pkg/Ai"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpGenerateAiRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	OpenaiService := Ai.NewOpenAiService()
	HistoryChatCareer := repositories.NewRecommendationAiRepository(db)
	UserRepo := repositories.NewUserRepository(db)
	RoleRepo := repositories.NewRoleRepository(db)
	CareerRepo := repositories.NewCareerRepository(db)

	UserService := services.NewUserService(services.UserServiceImpl{
		UserRepo:  UserRepo,
		Validator: validate,
		RoleRepo:  RoleRepo,
	})

	RecommendedAiService := services.NewRecommendedAiService(services.RecommendedAiServiceImpl{
		OpenAi:            OpenaiService,
		UserService:       UserService,
		CareerRepo:        CareerRepo,
		HistoryChatCareer: HistoryChatCareer,
		Validate:          *validate,
	})

	RecommendedAiHandler := handlers.NewRecommendedAiService(handlers.RecommendedAiHandlerImpl{
		RecommendedAiService: RecommendedAiService,
	})

	verifyUser := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	verifyUser.POST("/career/generate-recommendation-career", RecommendedAiHandler.GenerateRecommendationHandler)
}
