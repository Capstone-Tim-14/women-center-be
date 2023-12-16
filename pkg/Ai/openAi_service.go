package Ai

import (
	"context"
	"woman-center-be/internal/app/v1/models/domain"

	"github.com/labstack/echo/v4"
	"github.com/sashabaranov/go-openai"
	"github.com/spf13/viper"
)

type OpenAiService interface {
	EmbeddedPromptByDataCareer(careers []domain.Career) string
	GenerateMessage(ctx echo.Context, prompt string, question string) string
}

type OpenAiServiceImpl struct {
	OpenAi *openai.Client
}

func NewOpenAiService() OpenAiService {
	client := openai.NewClient(viper.GetString("OPENAI.TOKEN"))

	return &OpenAiServiceImpl{
		OpenAi: client,
	}
}

func (service *OpenAiServiceImpl) EmbeddedPromptByDataCareer(careers []domain.Career) string {

	var resultPrompting string

	if len(careers) > 0 {
		resultPrompting = "Sebagai narasumber yang handal, Anda ditugaskan untuk memberikan daftar rekomendasi jenjang karier yang sedang tren untuk wanita saat ini. Berikan informasi tentang posisi atau industri yang menjanjikan pertumbuhan karier yang pesat berdasarkan data-data berikut "

		for _, item := range careers {
			resultPrompting += item.Title_job + " " + item.About_job + " " + item.Location + " " + item.Linkedin_url + " "
		}

		resultPrompting += "lalu berikan hasil sesuai format-format dan link yang tertera pada data pada data-data yang sudah direkomendasikan."

	} else {
		resultPrompting = "Sebagai narasumber yang handal, Anda ditugaskan untuk memberikan daftar rekomendasi jenjang karier yang sedang tren untuk wanita saat ini. Berikan informasi tentang posisi atau industri yang menjanjikan pertumbuhan karier yang pesat, serta saran tentang keterampilan dan kualifikasi yang diperlukan. Pertimbangkan tren terkini dalam dunia kerja dan berikan wawasan mengenai bagaimana wanita dapat memanfaatkannya untuk memajukan karier mereka. Pastikan untuk memperhatikan diversifikasi industri dan cara mendukung kesetaraan gender di tempat kerja."
	}

	return resultPrompting
}

func (service *OpenAiServiceImpl) GenerateMessage(ctx echo.Context, prompt string, question string) string {

	result := make(chan string)

	go func() {
		client := service.OpenAi
		model := openai.GPT3Dot5Turbo
		messages := []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: prompt,
			}, {
				Role:    openai.ChatMessageRoleUser,
				Content: question,
			},
		}

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    model,
				Messages: messages,
			},
		)

		if err != nil {
			result <- "Error generating chat"
			return
		}

		chat := resp.Choices[0].Message.Content

		result <- chat
	}()

	select {
	case chatResult := <-result:
		return chatResult
	case <-ctx.Request().Context().Done():
		return "Generating cancelled"
	}

}
