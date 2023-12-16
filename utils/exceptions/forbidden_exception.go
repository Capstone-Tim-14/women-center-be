package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func StatusForbiddenResponse(ctx echo.Context, err error) error {
	return errorResponse(ctx, http.StatusForbidden, err.Error())
}
