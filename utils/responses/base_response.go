package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseField struct {
	Status  int
	Message string
	Data    interface{}
}

func BaseResponse(ctx echo.Context, status int, message string, data interface{}) error {
	return ctx.JSON(http.StatusOK, BaseField{
		Status:  http.StatusOK,
		Message: message,
		Data:    data,
	})
}
