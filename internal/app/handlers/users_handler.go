package handlers

import (
	"woman-center-be/internal/app/services"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	RegisterHandler(echo.Context) error
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(user services.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: user,
	}
}

func (user *UserHandlerImpl) RegisterHandler(ctx echo.Context) error {
	return nil
}
