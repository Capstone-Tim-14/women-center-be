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

func HttpCounselorRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	SpecialistRepo := repositories.NewSpecialistRepository(db)
	ScheduleRepo := repositories.NewScheduleRepository(db)
	AdminRepo := repositories.NewAdminRepository(db)
	CounselorHasSpecialistRepo := repositories.NewCounselorHasSpecialistRepository(db)
	CounselorService := services.NewCounselorService(services.CounselorServiceImpl{
		CounselorRepo:              CounselorRepo,
		RoleRepo:                   RoleRepo,
		Validator:                  validate,
		AdminRepo:                  AdminRepo,
		SpecialistRepo:             SpecialistRepo,
		CounselorHasSpecialistRepo: CounselorHasSpecialistRepo,
	})
	CounselorScheduleService := services.NewScheduleService(ScheduleRepo, validate, CounselorRepo)

	CounselorHandler := handlers.NewCounselorHandler(CounselorService)
	CounselorScheduleHandler := handlers.NewCounselorScheduleHandler(handlers.ScheduleHandlerImpl{
		CounselorService:         CounselorService,
		CounselorScheduleService: CounselorScheduleService,
	})

	userVerify := group.Group("/counselors", middlewares.VerifyTokenSignature("SECRET_KEY"))

	userVerify.GET("", CounselorHandler.GetCounselorsForMobile)
	userVerify.PUT("", CounselorHandler.UpdateCounselorForMobile)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	verifyTokenAdmin.POST("/counselor/:id/add-specialist", CounselorHandler.AddSpecialist)
	verifyTokenAdmin.DELETE("/counselor/:id/remove-specialist", CounselorHandler.RemoveManySpecialist)
	verifyTokenAdmin.POST("/counselors/register", CounselorHandler.RegisterHandler)
	verifyTokenAdmin.GET("/counselors", CounselorHandler.GetAllCounselorsHandler)
	verifyTokenAdmin.PUT("/counselors/:id", CounselorHandler.UpdateCounselorHandler)
	verifyTokenAdmin.POST("/counselor/:id/schedule", CounselorScheduleHandler.CreateScheduleHandler)
}
