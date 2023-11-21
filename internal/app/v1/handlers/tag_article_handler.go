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

type TagHandler interface {
	CreateTag(ctx echo.Context) error
	GetAllTagsHandler(ctx echo.Context) error
	DeleteTagByIdHandler(echo.Context) error
}

type TagHandlerImpl struct {
	TagService services.TagService
}

func NewTagHandler(tagService services.TagService) TagHandler {
	return &TagHandlerImpl{
		TagService: tagService,
	}
}

func (handler *TagHandlerImpl) CreateTag(ctx echo.Context) error {
	tagCreateRequest := requests.TagArticleRequest{}
	err := ctx.Bind(&tagCreateRequest)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	// Call TagService to create a new tag
	response, validation, err := handler.TagService.CreateTag(ctx, tagCreateRequest)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Error validation", validation)
	}

	if err != nil {

		if strings.Contains(err.Error(), "tag already exists") {
			return exceptions.StatusAlreadyExist(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	// Prepare the response
	tagCreateResponse := conversion.TagArticleDomainToTagArticleResponse(response)

	return responses.StatusCreated(ctx, "Tag created successfully", tagCreateResponse)
}

func (handler *TagHandlerImpl) GetAllTagsHandler(ctx echo.Context) error {
	response, err := handler.TagService.GetListTags(ctx)
	if err != nil {

		if strings.Contains(err.Error(), "Tags is empty") {
			return exceptions.StatusNotFound(ctx, err)
		}

		return exceptions.StatusInternalServerError(ctx, err)
	}

	tagResponse := conversion.ConvertTagResource(response)

	return responses.StatusOK(ctx, "Success Get All Tag", tagResponse)

}

func (handler *TagHandlerImpl) DeleteTagByIdHandler(ctx echo.Context) error {
	tagId := ctx.Param("id")
	tagIdInt, err := strconv.Atoi(tagId)
	if err != nil {
		return exceptions.StatusBadRequest(ctx, err)
	}

	err = handler.TagService.DeleteTagById(ctx, tagIdInt)
	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Success Delete Tag By Id", nil)

}
