package application

import (
	"context"
	"order-service/internal/domain"
	"order-service/internal/ports"
)

type CreateOrder struct {
	repo ports.OrderRepository
}

func NewCreateOrder(r ports.OrderRepository) *CreateOrder {
	return &CreateOrder{
		repo: r,
	}
}

// Execute handles the creation of a new order.
// IMPORTANT: This use case does NOT publish events directly.
// Instead, it relies on the Outbox Pattern to guarantee consistency
// between database state and event publishing.
func (uc *CreateOrder) Execute(ctx context.Context, order *domain.Order) error {
	return uc.repo.Save(ctx, order)
}
