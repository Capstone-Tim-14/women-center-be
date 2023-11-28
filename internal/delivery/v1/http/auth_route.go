package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpAuthRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	RoleRepo := repositories.NewRoleRepository(db)
	CounselorRepo := repositories.NewCounselorRepository(db)
	UserRepo := repositories.NewUserRepository(db)
	AdminRepo := repositories.NewAdminRepository(db)
	CredentialRepo := repositories.NewCredentialRepository(db)

	OTPService := services.NewOtpServiceImpl(services.OtpServiceImpl{
		UserRepo: UserRepo,
	})

	OTPHandler := handlers.NewOtpHandlerImpl(handlers.OtpHandlerImpl{
		OtpService: OTPService,
	})

	AuthService := services.NewAuthService(services.AuthServiceImpl{
		AdminRepo:      AdminRepo,
		RoleRepo:       RoleRepo,
		UserRepo:       UserRepo,
		CounselorRepo:  CounselorRepo,
		CredentialRepo: CredentialRepo,
		Validate:       validate,
	})
	AuthHandler := handlers.NewAuthHandler(AuthService)

	group.GET("/google-auth", AuthHandler.OauthGoogleHandler)
	group.GET("/callback-google-auth", AuthHandler.OauthCallbackGoogleHandler)
	group.POST("/auth/otp/generate", OTPHandler.SendOtpHandler)
	group.POST("/auth/otp/verify", OTPHandler.VerifyTokenHandler)

	Admin := group.Group("/admin")

	group.POST("/auth", AuthHandler.UserAuthHandler)
	Admin.POST("/auth", AuthHandler.AdminAuthHandler)

}
