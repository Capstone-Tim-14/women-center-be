package responses

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type WebResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Meta    any    `json:"meta,omitempty"`
	Data    any    `json:"data,omitempty"`
}

func successResponse(ctx echo.Context, code int, message string, data any) error {
	return ctx.JSON(code, WebResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func successResponseWithMeta(ctx echo.Context, code int, message string, meta any, data any) error {
	return ctx.JSON(code, WebResponse{
		Code:    code,
		Message: message,
		Meta:    meta,
		Data:    data,
	})
}

func StatusCreated(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusCreated, message, data)
}

func StatusOK(ctx echo.Context, message string, data any) error {
	return successResponse(ctx, http.StatusOK, message, data)
}

func StatusOKWithMeta(ctx echo.Context, message string, meta any, data any) error {
	return successResponseWithMeta(ctx, http.StatusOK, message, meta, data)
}
