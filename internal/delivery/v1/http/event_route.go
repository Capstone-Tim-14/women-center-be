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

func HttpEventRoute(group *echo.Group, db *gorm.DB, validate *validator.Validate) {

	EventRepo := repositories.NewEventRepository(db)
	EventService := services.NewEventService(services.EventServiceImpl{
		EventRepo: EventRepo,
		Validator: validate,
	})
	EventHandler := handlers.NewEventHandler(EventService)

	verifyTokenAdmin := group.Group("/admin", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	EventGroup := verifyTokenAdmin.Group("")

	EventGroup.POST("/event", EventHandler.CreateEvent)
	EventGroup.GET("/event/:id", EventHandler.GetDetailEvent)

	verifyTokenUser := group.Group("", middlewares.VerifyTokenSignature("SECRET_KEY"))

	verifyTokenUser.GET("/event/:id", EventHandler.GetDetailEventMobile)

}
