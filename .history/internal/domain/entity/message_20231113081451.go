package entity

import (
	"time"

	"github.com/google/uuid"
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
		ID:        uuid.New().String(),
		Role:      role,
		Content:   content,
		Model:     model,
		CreatedAt: time.Now(),
	}
	return msg, nil
}

func (m *Message) Validate() error {
	if m.Role != "user" && m.Role != "system" && m.Role != "assistant" {
		return errors.New("Invalid role")
	}
	if m.Content == "" {
		return errors.New("Content is empty")
	}
	if m.CreatedAt.IsZero() {
		return errors.New("Invalid created_at")
	}
	return nil
}
