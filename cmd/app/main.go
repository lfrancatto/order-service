package main

import (
	"log"
	"order-service/pkg/config"
)

func main() {

	cfg := config.Load()

	log.Println("Kafka:", cfg.KafkaBrokers)
	log.Println("Postgres:", cfg.PostgresDSN)
}
