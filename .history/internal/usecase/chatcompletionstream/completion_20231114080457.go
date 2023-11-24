package chatcompletionstream

import (
	"github.com/MrJulay/fclx/chatservice/internal/domain/gateway"
	openai "github.com\"
)

type ChatCompletionUseCase struct {
	ChatGateway gateway.ChatGateway
	OpenAiClient 
}
