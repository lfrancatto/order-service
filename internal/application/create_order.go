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

func (uc *CreateOrder) Execute(ctx context.Context, order *domain.Order) error {
	return uc.repo.Save(ctx, order)
}
