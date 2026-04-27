package main

import (
	"database/sql"
	"log"
	nethttp "net/http"

	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"order-service/internal/adapters/db"
	httpadapter "order-service/internal/adapters/http"
	"order-service/internal/adapters/kafka"
	"order-service/internal/application"
	"order-service/internal/consumers"
	"order-service/internal/infrastructure"
	"order-service/internal/services"
	"order-service/pkg/config"
)

func main() {
	cfg := config.Load()

	// DB connection
	dbConn, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	// Run migrations
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(dbConn, "migrations"); err != nil {
		log.Fatal(err)
	}

	// Dependencies
	repo := db.NewPostgresRepository(dbConn)
	usecase := application.NewCreateOrder(repo)
	handler := httpadapter.NewHandler(usecase)

	// Kafka producer (used by outbox + DLQ)
	producer, err := kafka.NewProducer(cfg.KafkaBrokers)
	if err != nil {
		log.Fatal(err)
	}

	// Outbox worker
	worker := infrastructure.NewOutboxWorker(dbConn, producer)
	go worker.Start()

	// Consumer setup
	processor := services.NewOrderProcessor()

	consumer, err := consumers.NewConsumer(
		cfg.KafkaBrokers,
		"order-group",
		processor.Process,
		producer,
	)
	if err != nil {
		log.Fatal(err)
	}

	go consumer.Start("orders")

	// HTTP server
	nethttp.HandleFunc("/orders", handler.CreateOrder)

	log.Println("Server running on :8081")
	nethttp.ListenAndServe(":8081", nil)
}
