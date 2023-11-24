package chatcompletionstream

import (
	"github.com/MrJulay/fclx/chatservice/internal/domain/gateway"

	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionConfigInputDTO struct {
	Model string
	ModelMaxToken int
	
}
type ChatCompletionUseCase struct {
	ChatGateway  gateway.ChatGateway
	OpenAiClient *openai.Client
}

func NewChatCompletionUseCase(chatGateway gateway.ChatGateway, openAiClient *openai.Client) *ChatCompletionUseCase {
	return &ChatCompletionUseCase{
		ChatGateway:  chatGateway,
		OpenAiClient: openAiClient,
	}
}
