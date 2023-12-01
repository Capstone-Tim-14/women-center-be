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

func HttpCounselingPackageRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {
	CounselingRepo := repositories.NewCounselingRepository(db)
	AdminRepo := repositories.NewAdminRepository(db)
	CounselingService := services.NewCounselingPackageService(services.CounselingPackageServiceImpl{
		CounselingPackageRepo: CounselingRepo,
		Validator:             validate,
		AdminRepo:             AdminRepo,
	})
	CounselingHandler := handlers.NewCounselingPackageHandler(CounselingService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))
	verifyTokenAdmin.GET("/counseling-packages", CounselingHandler.GetAllPackage)
	//verifyTokenAdmin.POST()
}
