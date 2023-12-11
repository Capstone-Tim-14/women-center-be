package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/delivery/v1/middlewares"
	payment "woman-center-be/pkg/payment/midtrans"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpBookingCounselingRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	UserRepo := repositories.NewUserRepository(db)
	RoleRepo := repositories.NewRoleRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	UserSchedule := repositories.NewUserScheduleConselingRepository(db)
	PackageRepo := repositories.NewCounselingPackageRepository(db)
	BookingRepo := repositories.NewBookingCounselingRepository(db)

	UserService := services.NewUserService(UserRepo, RoleRepo, validate)
	MidtransService := payment.NewMidtransCoreApiImpl()

	TransactionService := services.NewTransactionPaymentService(services.TransactionPaymentServiceImpl{
		BookingRepo:     BookingRepo,
		MidtransService: MidtransService,
	})

	CounselingPackageService := services.NewCounselingPackageService(services.CounselingPackageServiceImpl{
		CounselingPackageRepo: PackageRepo,
	})

	BookingCounselingService := services.NewBookingService(services.BookingServiceImpl{
		UserService:              UserService,
		UserRepo:                 UserRepo,
		ScheduleRepo:             UserSchedule,
		CounselorRepo:            CounselorRepo,
		BookingRepo:              BookingRepo,
		CounselingPackage:        PackageRepo,
		CounselingPackageService: CounselingPackageService,
		Validate:                 validate,
	})

	BookingCounselingHandler := handlers.NewBookingCounselingHandler(handlers.BookingCounselingHandlerImpl{
		BookingService:     BookingCounselingService,
		TransactionService: TransactionService,
	})

	verifyToken := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	verifyToken.POST("/booking", BookingCounselingHandler.CreateBookingHandler)
	verifyToken.POST("/charge-payment", BookingCounselingHandler.CreateTransactionPaymentHandler)

}
