package entity

type Message struct{
	ID string
	Role string
	Content string
	Tokens int
	Model *Model
	CreatedAt time.Time
}
