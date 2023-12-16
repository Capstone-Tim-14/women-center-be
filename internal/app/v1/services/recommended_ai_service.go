package services

import (
	"fmt"
	"woman-center-be/internal/app/v1/repositories"
	conversion "woman-center-be/internal/web/conversion/request/v1"
	conRes "woman-center-be/internal/web/conversion/resource/v1"
	"woman-center-be/internal/web/requests/v1"
	"woman-center-be/internal/web/resources/v1"
	"woman-center-be/pkg/Ai"
	"woman-center-be/utils/exceptions"
	"woman-center-be/utils/helpers"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type RecommendedAiService interface {
	GetRecommendedCareers(ctx echo.Context, message string) string
	SaveGenerateRecommendationCareer(ctx echo.Context, request requests.ChatRecomendedRequest) (*resources.GenerateRecommendationCareerResource, []exceptions.ValidationMessage, error)
	GetAllHistoryRecommendationCareer(ctx echo.Context) (*resources.HistoryRecommendationCareerAiResource, error)
}

type RecommendedAiServiceImpl struct {
	OpenAi            Ai.OpenAiService
	UserService       UserService
	CareerRepo        repositories.CareerRepository
	HistoryChatCareer repositories.RecommendationAiRepository
	Validate          validator.Validate
}

func NewRecommendedAiService(recommendedService RecommendedAiServiceImpl) RecommendedAiService {
	return &recommendedService
}

func (service *RecommendedAiServiceImpl) SaveGenerateRecommendationCareer(ctx echo.Context, request requests.ChatRecomendedRequest) (*resources.GenerateRecommendationCareerResource, []exceptions.ValidationMessage, error) {

	validation := service.Validate.Struct(request)

	if validation != nil {
		return nil, helpers.ValidationError(ctx, validation), nil
	}

	UserProfile, errClaims := service.UserService.GetUserProfile(ctx)

	if errClaims != nil {
		return nil, nil, fmt.Errorf("Error claim your user, please authentication")
	}

	GenerateMessage := service.GetRecommendedCareers(ctx, request.Message)

	if GenerateMessage == "Error generating chat" || GenerateMessage == "Generating cancelled" {
		return nil, nil, fmt.Errorf(GenerateMessage)
	}

	ConversionChat := conversion.ConvertChatToHistoryRecommendationCareer(*UserProfile, request.Message, GenerateMessage)

	SaveChat, errSaveChat := service.HistoryChatCareer.SaveHistoryRecommendationCareer(ConversionChat)

	if errSaveChat != nil {
		return nil, nil, errSaveChat
	}

	Response := conRes.ConvertSaveHistoryChatToCareerRecommendationResource(*UserProfile, *SaveChat)

	return &Response, nil, nil
}

func (service *RecommendedAiServiceImpl) GetRecommendedCareers(ctx echo.Context, message string) string {

	GetCareerList, _ := service.CareerRepo.GetAllCareerNoFilter()
	SetPromptingCareer := service.OpenAi.EmbeddedPromptByDataCareer(GetCareerList)

	GenerateMessage := service.OpenAi.GenerateMessage(ctx, SetPromptingCareer, message)

	return GenerateMessage
}

func (service *RecommendedAiServiceImpl) GetAllHistoryRecommendationCareer(ctx echo.Context) (*resources.HistoryRecommendationCareerAiResource, error) {
	UserProfile, errClaims := service.UserService.GetUserProfile(ctx)

	if errClaims != nil {
		return nil, fmt.Errorf("Error claim your user, please authentication")
	}

	_, errGetHistory := service.HistoryChatCareer.FindAllHistoryRecommendationCareer()

	if errGetHistory != nil {
		return nil, fmt.Errorf("Error get history")
	}

	ConversionHistory := conRes.ConvertHistoryChatToCareerRecommendationResource(UserProfile)

	return &ConversionHistory, nil
}
