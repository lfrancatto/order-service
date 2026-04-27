package domain

import "time"

type Order struct {
	ID        string
	UserID    string
	Amount    float64
	CreatedAt time.Time
}

func NewOrder(id, userID string, amount float64) *Order {
	return &Order{
		ID:        id,
		UserID:    userID,
		Amount:    amount,
		CreatedAt: time.Now(),
	}
}
