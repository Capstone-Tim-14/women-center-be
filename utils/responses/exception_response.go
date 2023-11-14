package responses

import "github.com/labstack/echo/v4"

type ExceptionField struct {
	Status  int
	Message string
}

func ExceptionResponse(ctx echo.Context, status int, message string) error {
	return ctx.JSON(status, ExceptionField{
		Status:  status,
		Message: message,
	})
}
