SERVICE := postgres pgadmin
DB_NAME := tutu
USER := postgres
PASSWORD := postgres

HOST := localhost
PORT := 5432

GOOSE_DRIVER := postgres
DB_STRING := "postgres://$(USER):$(PASSWORD)@$(HOST):$(PORT)/$(DB_NAME)?sslmode=disable"
MIGRATIONS := ./migrations

all: compose-up goose-up

# docker compose
compose-up:
	@docker-compose up -d $(SERVICE)

compose-down:
	@docker-compose down $(SERVICE)

compose-stop:
	@docker-compose stop $(SERVICE)

compose-start:
	@docker-compose start $(SERVICE)

compose-ps:
	@docker-compose ps $(SERVICE)

compose-rm:
	@docker-compose rm $(SERVICE)


# migrations
goose-install:
	go install github.com/pressly/goose/v3/cmd/goose@latest

goose-add:
	@goose -dir $(MIGRATIONS) $(GOOSE_DRIVER) $(DB_STRING) create rename_me sql

goose-up:
	@goose -dir $(MIGRATIONS) $(GOOSE_DRIVER) $(DB_STRING) up

goose-down:
	@goose -dir $(MIGRATIONS) $(GOOSE_DRIVER) $(DB_STRING) down

goose-down-all:
	@goose -dir $(MIGRATIONS) $(GOOSE_DRIVER) $(DB_STRING) down-to 0

goose-restart: goose-down-all goose-up
	cd generator && go run main.go

goose-status:
	@goose -dir $(MIGRATIONS) $(GOOSE_DRIVER) $(DB_STRING) status

