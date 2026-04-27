.PHONY: env-up env-down

env-up:
	docker compose up -d

env-down:
	docker compose down
