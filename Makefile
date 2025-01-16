-include .env
.PHONY: docs

build:
	@echo "Building..."
	@go build -o bin/main cmd/api/main.go

clean:
	@echo "Cleaning"
	@rm -f bin/main

start:
	@bin/main

run:
	@go run cmd/api/main.go

config:
	@go run cmd/config/main.go

watch:
	@air

lint:
	@go vet ./...

docs:
	@go mod vendor
	@touch docs/swagger.json
	@swag init -o docs -d cmd/api,internal,vendor/github.com/dmytro-kucherenko/smartner-utils-package
	@swag fmt

pre-commit:
	@pre-commit autoupdate && pre-commit install

db-up:
	@docker-compose up -d

db-down:
	@docker-compose down

db-start:
	@docker-compose start

db-stop:
	@docker-compose stop

db-gen:
	@sqlc generate

migration-create:
	@migrate create -ext sql -dir internal/db/migrations $(name)

migration-up:
	@migrate -path internal/db/migrations -database "${DB_CONNECTION}" up

migration-down:
	@migrate -path internal/db/migrations -database "${DB_CONNECTION}" down 1

migration-reset:
	@migrate -path internal/db/migrations -database "${DB_CONNECTION}" down
