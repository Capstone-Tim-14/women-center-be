package middlewares

import (
	"errors"
	"net/http"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func VerifyTokenSignature(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(viper.GetString(secretKey)),
		ErrorHandler: func(ctx echo.Context, err error) error {
			return echo.NewHTTPError(http.StatusUnauthorized, exceptions.StatusUnauthorizedResponse(ctx, errors.New("Unauthorized!")))
		},
		ContextKey: "auth-login",
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.JwtProfileClaims)
		},
	})
}
