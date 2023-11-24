package entity

import (
	"time"

	"github.com/google"
)

type Message struct {
	ID        string
	Role      string
	Content   string
	Tokens    int
	Model     *Model
	CreatedAt time.Time
}

func NewMessage(role, content string, model *Model) (*Message, error) {
	msg := &Message{
		ID: uuid.New().String(),
	}
}
