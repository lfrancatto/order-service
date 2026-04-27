package main

import (
	"database/sql"
	"log"
	nethttp "net/http"
	"order-service/internal/adapters/db"
	httpadapter "order-service/internal/adapters/http"
	"order-service/internal/adapters/kafka"
	"order-service/internal/application"
	"order-service/internal/infrastructure"
	"order-service/pkg/config"

	"github.com/pressly/goose/v3"
)

func main() {

	cfg := config.Load()

	dbConn, err := sql.Open("postgres", cfg.PostgresDSN)
	if err != nil {
		log.Fatal(err)
	}

	//Goose setup
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(dbConn, "migrations"); err != nil {
		log.Fatal(err)
	}

	repo := db.NewPostgresRepository(dbConn)
	usecase := application.NewCreateOrder(repo)

	handler := httpadapter.NewHandler(usecase)

	producer, _ := kafka.NewProducer(cfg.KafkaBrokers)

	worker := infrastructure.NewOutboxWorker(dbConn, producer)
	go worker.Start()

	nethttp.HandleFunc("/orders", handler.CreateOrder)

	log.Println("Server running on :8081")
	nethttp.ListenAndServe(":8081", nil)
}
