package v1

import (
	"net/http"
	routes "woman-center-be/internal/delivery/v1/http"
	"woman-center-be/internal/delivery/v1/middlewares"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitRoutes(e *echo.Echo, db *gorm.DB, validator *validator.Validate) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	v1 := e.Group("/api/v1", middlewares.Logger())

	VerifyTokenAdmin := v1.Group("", middlewares.VerifyTokenSignature("SECRET_KEY_ADMIN"))

	routes.HttpUserRoute(v1, db, validator)
	routes.HttpAuthRoute(v1, db, validator)
	routes.HttpRoleRoute(VerifyTokenAdmin, db, validator)
	routes.HttpCounselorRoute(v1, db, validator)
	routes.HttpAdminRoute(v1, db, validator)
	routes.HttpTagRoute(v1, db, validator)
	routes.HttpArticleRoute(v1, db, validator)
	routes.HttpSpecialistRoute(v1, db, validator)
	routes.HttpCareerRoute(v1, db, validator)
	routes.HttpJobTypeRoute(v1, db, validator)

	v1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome to woman center api")
	})

}
