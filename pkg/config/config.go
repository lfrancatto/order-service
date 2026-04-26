package config

type Config struct {
	KafkaBrokers string
	PostgresDSN  string
}	

func Load() *Config {
	return &Config{
		KafkaBrokers: "localhost:9092",
		PostgresDSN:  "postgres://postgres:postgres@localhost:5432/orders?sslmode=disable",
	}
}
