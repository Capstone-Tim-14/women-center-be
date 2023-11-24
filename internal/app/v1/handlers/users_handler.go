package handlers

import (
	"strings"
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	RegisterHandler(echo.Context) error
	ProfileHandler(echo.Context) error
	UpdateProfileHandler(echo.Context) error
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(user services.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: user,
	}
}

func (h *UserHandlerImpl) ProfileHandler(ctx echo.Context) error {
	getProfileUser, err := h.UserService.GetUserProfile(ctx)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	conversionProfile := conversion.UserDomainToUserProfileResource(getProfileUser)

	return responses.StatusOK(ctx, "Success get user profile", conversionProfile)
}

func (handler *UserHandlerImpl) RegisterHandler(ctx echo.Context) error {
	userCreateRequest := requests.UserRequest{}
	err := ctx.Bind(&userCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.UserService.RegisterUser(ctx, userCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	userCreateResponse := conversion.UserDomainToUserResource(response)

	return responses.StatusCreated(ctx, "User created successfully", userCreateResponse)

}

func (handler *UserHandlerImpl) UpdateProfileHandler(ctx echo.Context) error {
	userUpdateRequest := requests.UpdateUserProfileRequest{}
	err := ctx.Bind(&userUpdateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	response, validation, err := handler.UserService.UpdateUserProfile(ctx, userUpdateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	// pake ini dulu biar tau respon di postman, ntar apus + ignore response
	userUpdateResponse := conversion.UserDomainToUserUpdateProfileResource(response)

	return responses.StatusOK(ctx, "User profile updated", userUpdateResponse)

}
