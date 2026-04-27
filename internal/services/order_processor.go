package services

import (
	"encoding/json"
	"errors"
)

// Order represents incoming event structure
type Order struct {
	ID     string  `json:"id"`
	UserID string  `json:"user_id"`
	Amount float64 `json:"amount"`
}

// OrderProcessor handles business logic for consumed messages
type OrderProcessor struct{}

func NewOrderProcessor() *OrderProcessor {
	return &OrderProcessor{}
}

// Process executes business logic for an order event
// IMPORTANT: This function must be idempotent in real systems
func (p *OrderProcessor) Process(msg []byte) error {
	var order Order

	err := json.Unmarshal(msg, &order)
	if err != nil {
		return err
	}

	// Simulate business rule failure
	if order.Amount > 1000 {
		return errors.New("amount exceeds allowed limit")
	}

	return nil
}