package entity

type ChatConfig struct {
	
}
type Chat struct {
	ID         string
	UserID     string
	Status     string
	TokenUsage int
	Config *ChatConfig
}
