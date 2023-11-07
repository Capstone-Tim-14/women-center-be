package v1

import (
	routes "woman-center-be/internal/delivery/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, validator *validator.Validate) {

	v1 := e.Group("/api/v1")

	routes.HttpUserRoute(v1, db, validator)

}
