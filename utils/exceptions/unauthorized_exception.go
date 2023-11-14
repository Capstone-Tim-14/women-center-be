package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StatusUnauthorizedResponse(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusUnauthorized, err.Error())
}
