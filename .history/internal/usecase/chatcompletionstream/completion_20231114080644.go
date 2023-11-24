package chatcompletionstream

import (
	"github.com/MrJulay/fclx/chatservice/internal/domain/gateway"
	openai "github.com/sashabaranov/go-openai"
)

type ChatCompletionUseCase struct {
	ChatGateway gateway.ChatGateway
	OpenAiClient *openai.Client
}

func NewChatCompletionUseCase (chatGateway gateway.ChatGateway, openAiClient Op)
