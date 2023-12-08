package handlers

import (
	"strconv"
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
	AddCounselorFavorite(echo.Context) error
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
	pictureProfile, _ := ctx.FormFile("picture_profile")

	err := ctx.Bind(&userUpdateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := handler.UserService.UpdateUserProfile(ctx, userUpdateRequest, pictureProfile)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		if strings.Contains(err.Error(), "Email already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "User profile updated", nil)

}

func (handler *UserHandlerImpl) AddCounselorFavorite(ctx echo.Context) error {

	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)
	var request requests.CounselorFavotireRequest
	errBinding := ctx.Bind(&request)

	if errBinding != nil {
		return exceptions.StatusBadRequest(ctx, errBinding)
	}

	validation, err := handler.UserService.AddCounselorFavorite(ctx, convertid, request)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success add user's counselor favorite", nil)

}
