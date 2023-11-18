package handlers

import (
	"woman-center-be/internal/app/v1/services"
	conversion "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type TagHandler interface {
	CreateTag(ctx echo.Context) error
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
		// Handle specific errors if needed
		return exceptions.StatusInternalServerError(ctx, err)
	}

	// Prepare the response
	tagCreateResponse := conversion.TagArticleDomainToTagArticleResponse(response)

	return responses.StatusCreated(ctx, "Tag created successfully", tagCreateResponse)
}

// fix harusnya
