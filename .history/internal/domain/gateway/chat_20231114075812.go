package gateway

import (
	"context"

	"github.com/MrJulay/fclx/chatservice/internal/domain/entity"
)

type ChatGateway interface {
	CreateChat(ctx context.Context, char *entity.Chat ) error
}