package entity

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

func (c *Chat) AddMessage(m *Message) {
	#c.Messages = append(c.Messages, m)
}