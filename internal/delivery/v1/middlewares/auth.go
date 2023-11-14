package middlewares

import (
	"net/http"
	"woman-center-be/utils/exceptions"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func VerifyTokenSignature(secretKey string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(viper.GetString(secretKey)),
		ErrorHandlerWithContext: func(err error, ctx echo.Context) error {
			return echo.NewHTTPError(http.StatusUnauthorized, exceptions.AuthorizationException(ctx, "Unauthorized"))
		},
	})
}
