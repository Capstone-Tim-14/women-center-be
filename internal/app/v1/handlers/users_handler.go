package handlers

import (
	"fmt"
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
	ListUserHandler(echo.Context) error
	UserDetailHandler(echo.Context) error
	UserUpdateHandler(echo.Context) error
	RegisterHandler(echo.Context) error
	ProfileHandler(echo.Context) error
	UpdateProfileHandler(echo.Context) error
	AddFavoriteArticleHandler(ctx echo.Context) error
	DeleteFavoriteArticleHandler(ctx echo.Context) error
	AllFavoriteArticleHandler(ctx echo.Context) error
	AddCounselorFavorite(echo.Context) error
	RemoveCounselorFavorite(echo.Context) error
	GetCounselorFavorite(echo.Context) error
}

type UserHandlerImpl struct {
	UserService services.UserService
}

func NewUserHandler(user services.UserService) UserHandler {
	return &UserHandlerImpl{
		UserService: user,
	}
}

func (h *UserHandlerImpl) UserUpdateHandler(ctx echo.Context) error {

	userUpdateRequest := requests.UpdateUserProfileRequest{}
	pictureProfile, _ := ctx.FormFile("picture_profile")

	err := ctx.Bind(&userUpdateRequest)

	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	_, validation, err := h.UserService.UpdateUser(ctx, userUpdateRequest, pictureProfile)

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

func (h *UserHandlerImpl) UserDetailHandler(ctx echo.Context) error {

	GetId := ctx.Param("id")
	ConvertId, errId := strconv.Atoi(GetId)

	if errId != nil {
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Invalid format id"))
	}

	UserRes, errGetUser := h.UserService.UserDetail(uint(ConvertId))

	if errGetUser != nil {
		return exceptions.StatusInternalServerError(ctx, errGetUser)
	}

	return responses.StatusOK(ctx, "Success get user detail", UserRes)
}

func (h *UserHandlerImpl) ListUserHandler(ctx echo.Context) error {

	GetUsers, errGetUser := h.UserService.UsersList()

	if errGetUser != nil {
		return exceptions.StatusBadRequest(ctx, errGetUser)
	}

	return responses.StatusOK(ctx, "Succcess get users", GetUsers)
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

func (h *UserHandlerImpl) AddFavoriteArticleHandler(ctx echo.Context) error {
	slug := ctx.Param("slug")

	err := h.UserService.AddFavoriteArticle(ctx, slug)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to find article") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success add article favorite", nil)
}

func (h *UserHandlerImpl) DeleteFavoriteArticleHandler(ctx echo.Context) error {
	slugArticle := ctx.Param("slug")

	err := h.UserService.DeleteFavoriteArticle(ctx, slugArticle)
	if err != nil {
		if strings.Contains(err.Error(), "Failed to find article") {
			return exceptions.StatusNotFound(ctx, err)
		}
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success remove article favorite", nil)
}

func (h *UserHandlerImpl) AllFavoriteArticleHandler(ctx echo.Context) error {
	user, err := h.UserService.AllFavoriteArticle(ctx)
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	userFavoriteResponse := conversion.UserFavoriteArticleResponse(user)

	return responses.StatusOK(ctx, "Success get all article favorite", userFavoriteResponse)
}

func (handler *UserHandlerImpl) AddCounselorFavorite(ctx echo.Context) error {

	id := ctx.Param("id")
	convertid, _ := strconv.Atoi(id)

	err := handler.UserService.AddCounselorFavorite(ctx, convertid)

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Success Add Counselor Favorite", nil)

}

func (handler *UserHandlerImpl) RemoveCounselorFavorite(ctx echo.Context) error {
	GetId := ctx.Param("id")

	ConvertId, errConvertId := strconv.Atoi(GetId)

	if errConvertId != nil {
		fmt.Errorf(errConvertId.Error())
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Invalid Format id"))
	}

	Validation, errGetUser := handler.UserService.RemoveCounselorFavorite(ctx, ConvertId)

	if Validation != nil {
		return exceptions.ValidationException(ctx, "Validation Error", Validation)
	}

	if errGetUser != nil {
		if strings.Contains(errGetUser.Error(), "User not found") {
			return exceptions.StatusNotFound(ctx, errGetUser)
		}
		if strings.Contains(errGetUser.Error(), "Counselor not found") {
			return exceptions.StatusNotFound(ctx, errGetUser)
		}

		return exceptions.StatusInternalServerError(ctx, errGetUser)
	}

	return responses.StatusOK(ctx, "Success Remove Counselor Favorite", nil)
}

func (handler *UserHandlerImpl) GetCounselorFavorite(ctx echo.Context) error {
	user, err := handler.UserService.GetCounselorFavorite(ctx)
	if err != nil {
		return exceptions.StatusNotFound(ctx, err)
	}

	userFavoriteResponse := conversion.UserCounselorFavoriteResponse(user)

	return responses.StatusOK(ctx, "Success Get All Counselor Favorite", userFavoriteResponse)
}
