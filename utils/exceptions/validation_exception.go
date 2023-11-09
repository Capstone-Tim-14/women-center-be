package exceptions

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type ValidationMessage struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func ValidationException(ctx echo.Context, message string, form []ValidationMessage) error {
	return ctx.JSON(http.StatusUnprocessableEntity, ErrorField{
		Code:    http.StatusUnprocessableEntity,
		Message: message,
		Errors:  form,
	})
}
