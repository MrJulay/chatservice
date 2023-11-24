package entity

type ChatConfig struct {
	Model *Model
	Temperature float32
	Topp float32
	N int
	stop []string
	
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
