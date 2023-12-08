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

func HttpBookingCounselingRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	UserRepo := repositories.NewUserRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	UserSchedule := repositories.NewUserScheduleConselingRepository(db)
	PackageRepo := repositories.NewCounselingPackageRepository(db)

	BookingCounselingService := services.NewBookingService(services.BookingServiceImpl{
		UserRepo:          UserRepo,
		ScheduleRepo:      UserSchedule,
		CounselorRepo:     CounselorRepo,
		CounselingPackage: PackageRepo,
		Validate:          validate,
	})

	BookingCounselingHandler := handlers.NewBookingCounselingHandler(handlers.BookingCounselingHandlerImpl{
		BookingService: BookingCounselingService,
	})

	verifyToken := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	verifyToken.POST("/booking", BookingCounselingHandler.CreateBookingHandler)

}
