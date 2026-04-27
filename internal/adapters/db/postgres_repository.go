package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"order-service/internal/domain"
)

const (
	insertOrder = `INSERT INTO orders (id, user_id, amount, created_at) VALUES ($1, $2, $3, $4)`

	insertOutbox = `INSERT INTO outbox (topic, key, payload) VALUES ($1, $2, $3)`
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(db *sql.DB) *PostgresRepository {
	return &PostgresRepository{
		db: db,
	}
}

func (r *PostgresRepository) Save(ctx context.Context, order *domain.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, insertOrder, order.ID, order.UserID, order.Amount, order.CreatedAt)
	if err != nil {
		tx.Rollback()
		return err
	}

	payload, _ := json.Marshal(order)
	_, err = tx.ExecContext(ctx, insertOutbox, "orders", order.UserID, payload)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
