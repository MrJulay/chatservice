package entity

import "errors"

type ChatConfig struct {
	Model            *Model
	Temperature      float32
	Topp             float32
	N                int
	stop             []string
	MaxTokens        int
	PresencePenalty  float32
	FrequencyPenalty float32
}

type Chat struct {
	ID                   string
	UserID               string
	InitialSystemMessage *Message
	Messages             []*Message
	ErasedMessages       []*Message
	Status               string
	TokenUsage           int
	Config               *ChatConfig
}

func (c *Chat) AddMessage(m *Message) error {
	if c.Status == "ended" {
		return errors.New("chat is ended. no more messages allowed")
	}
	for {
		if c.Config.Model.GetMaxTokens() >= m.GetQtdTokens() + c.TokenUsage {
		//c.Messages = append(c.Messages, m)
		}
	}

}

func (c *Chat) RefreshTokenUsage() {
	c.TokenUsage = 0
	for m := range c.Messages {
		c.TokenUsage += c.Messages[m].GetQtdTokens()
	}
}
