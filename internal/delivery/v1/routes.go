package v1

import (
	"net/http"
	routes "woman-center-be/internal/delivery/v1/http"
	"woman-center-be/internal/delivery/v1/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, validator *validator.Validate) {

	v1 := e.Group("/api/v1", middlewares.Logger())

	routes.HttpUserRoute(v1, db, validator)
	routes.HttpRoleRoute(v1, db, validator)

	routes.HttpCounselorRoute(v1, db, validator)

	v1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to woman center api")
	})

}
