package application

import (
	"order-service/internal/domain"
	"testing"
)

// Mock repository to isolate the use case
type mockRepo struct {
	called bool
}

func (m *mockRepo) Save(order *domain.Order) error {
	m.called = true
	return nil
}

func TestCreateOrderExecute(t *testing.T) {
	repo := &mockRepo{}
	usecase := NewCreateOrder(repo)

	order := domain.NewOrder("1", "u1", 50)

	err := usecase.Execute(order)

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !repo.called {
		t.Error("expected repository Save to be called")
	}
}
