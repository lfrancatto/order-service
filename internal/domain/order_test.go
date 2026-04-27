package domain

import "testing"

func TestNewOrder(t *testing.T) {
	order := NewOrder("1", "user-1", 100.0)

	if order.ID != "1" {
		t.Errorf("expected ID to be 1, got %s", order.ID)
	}

	if order.UserID != "user-1" {
		t.Errorf("expected UserID to be user-1, got %s", order.UserID)
	}

	if order.Amount != 100.0 {
		t.Errorf("expected Amount to be 100, got %f", order.Amount)
	}

	if order.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}
}
