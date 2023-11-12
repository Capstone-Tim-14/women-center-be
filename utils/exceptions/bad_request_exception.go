package exceptions

import (
	"net/http"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

func BadRequestException(message string, ctx echo.Context) error {
	return responses.ExceptionResponse(ctx, http.StatusBadRequest, message)
}
