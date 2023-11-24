package chatcompletionstream

import (
	openai 
	"github.com/MrJulay/fclx/chatservice/internal/domain/gateway"
)

type ChatCompletionUseCase struct {
	ChatGateway gateway.ChatGateway
	OpenAiClient 
}
