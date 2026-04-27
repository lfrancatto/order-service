package infrastructure

import (
	"database/sql"
	"log"
	"order-service/internal/ports"
	"time"
)

type OutboxWorker struct {
	db        *sql.DB
	publisher ports.EventPublisher
}

func NewOutboxWorker(db *sql.DB, publisher ports.EventPublisher) *OutboxWorker {
	return &OutboxWorker{
		db:        db,
		publisher: publisher,
	}
}

func (w *OutboxWorker) Start() {
	for {
		rows, err := w.db.Query(`
			SELECT id, topic, key, payload
			FROM outbox
			WHERE processed = false
			LIMIT 10
		`)

		if err != nil {
			log.Println(err)
			continue
		}

		for rows.Next() {
			var (
				id         int
				topic, key string
				payload    []byte
			)

			rows.Scan(&id, &topic, &key, &payload)

			err = w.publisher.Publish(topic, key, payload)
			if err == nil {
				w.db.Exec("UPDATE outbox SET processed = true WHERE id = $1", id)
			}
		}

		time.Sleep(2 * time.Second)
	}
}
