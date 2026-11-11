package domain

type Product struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Price       float64   `json:"price"`
	Embedding   []float32 `json:"embedding,omitempty"`
}

type UserEvent struct {
	EventID   string `json:"eventId"`
	UserID    string `json:"userId"`
	Action    string `json:"action"` // view, click, purchase
	ProductID string `json:"productId"`
	Timestamp int64  `json:"timestamp"`
}
