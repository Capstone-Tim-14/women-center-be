package helpers

import (
	"fmt"
	"time"
	"woman-center-be/internal/web/resources/v1"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type JwtProfileClaims struct {
	Id       uint   `json:"id"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func GenerateToken(data resources.AuthResource, ctx echo.Context) (string, error) {

	var SetJWTProfile JwtProfileClaims

	SetJWTProfile.Id = data.Id
	SetJWTProfile.FullName = data.Fullname
	SetJWTProfile.Role = data.Role
	SetJWTProfile.Email = data.Email
	SetJWTProfile.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * 12))

	SigningToken := jwt.NewWithClaims(jwt.SigningMethodHS256, SetJWTProfile)

	var token string
	var errToken error

	if data.Role == "admin" || data.Role == "super admin" {
		token, errToken = SigningToken.SignedString([]byte(viper.GetString("SECRET_KEY_ADMIN")))
	} else if data.Role == "user" || data.Role == "counselor" {
		token, errToken = SigningToken.SignedString([]byte(viper.GetString("SECRET_KEY")))
	} else {
		return "", fmt.Errorf("Failed generate token")
	}

	if errToken != nil {
		return "", errToken
	}

	return token, nil

}
