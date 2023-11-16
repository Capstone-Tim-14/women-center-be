package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetAuthClaims(ctx echo.Context) *JwtProfileClaims {

	GetUserLogin := ctx.Get("auth-login").(*jwt.Token)
	ClaimUser := GetUserLogin.Claims.(*JwtProfileClaims)

	return ClaimUser
}
