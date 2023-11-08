package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

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

func (handler *UserHandlerImpl) RegisterHandler(ctx echo.Context) error {
	userCreateRequest := requests.UserRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, err := handler.UserService.RegisterUser(ctx, userCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "Validation failed") {
			return exceptions.StatusBadRequest(ctx, err)
		}

		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "User created successfully", response)

}
