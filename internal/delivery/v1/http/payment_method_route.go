package http

import (
	"woman-center-be/internal/app/v1/handlers"
	"woman-center-be/internal/app/v1/repositories"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/delivery/v1/middlewares"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func HttpPaymentMethodRoute(group *echo.Group, db *gorm.DB) {

	PaymentMethodRepo := repositories.NewPaymentMethodRepository(db)
	PaymentMethodService := services.NewPaymentMethodService(services.PaymentMethodServiceImpl{
		PaymentMethodRepo: PaymentMethodRepo,
	})
	PaymentMethodHandler := handlers.NewPaymentMethodHandler(handlers.PaymentMethodHandlerImpl{
		PaymentMethodService: PaymentMethodService,
	})

	verifyToken := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))
	verifyToken.GET("/payment-methods", PaymentMethodHandler.ListPaymentMethodHandler)
}
