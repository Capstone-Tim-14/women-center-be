package handlers

import (
	"fmt"
	"woman-center-be/internal/app/v1/services"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/responses"

	"github.com/labstack/echo/v4"
)

type RecommendedAiHandler interface {
	GenerateRecommendationHandler(echo.Context) error
	GetAllHistoryChatHandler(echo.Context) error
}

type RecommendedAiHandlerImpl struct {
	RecommendedAiService services.RecommendedAiService
}

func NewRecommendedAiService(recommendedAiHandler RecommendedAiHandlerImpl) RecommendedAiHandler {
	return &recommendedAiHandler
}

func (handler *RecommendedAiHandlerImpl) GenerateRecommendationHandler(ctx echo.Context) error {

	var request requests.ChatRecomendedRequest

	bindingReq := ctx.Bind(&request)

	if bindingReq != nil {
		return exceptions.StatusBadRequest(ctx, fmt.Errorf("Invalid format request"))
	}

	Response, validation, err := handler.RecommendedAiService.SaveGenerateRecommendationCareer(ctx, request)

	if validation != nil {
		return exceptions.ValidationException(ctx, "Validation request", validation)
	}

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusCreated(ctx, "Answering your question", Response)

}

func (handler *RecommendedAiHandlerImpl) GetAllHistoryChatHandler(ctx echo.Context) error {

	Response, err := handler.RecommendedAiService.GetAllHistoryRecommendationCareer(ctx)

	if err != nil {
		return exceptions.StatusInternalServerError(ctx, err)
	}

	return responses.StatusOK(ctx, "Get all history chat", Response)

}
