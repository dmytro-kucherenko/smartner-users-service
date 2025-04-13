-include .env
.PHONY: docs

build:
	@echo "Building..."
	@go build -o bin/api cmd/api/main.go

clean:
	@echo "Cleaning"
	@rm -f bin/api

start:
	@bin/api

run:
	@go run cmd/api/main.go

watch:
	@air -c api.air.toml

lint:
	@go vet ./...

docs:
	@go mod vendor
	@touch docs/swagger.json
	@go tool github.com/swaggo/swag/cmd/swag init -o docs -d cmd/api,internal,vendor/github.com/dmytro-kucherenko/smartner-utils-package
	@go tool github.com/swaggo/swag/cmd/swag fmt

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
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest create -ext sql -dir db/migrations $(name)

migration-up:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}" up

migration-down:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}" down 1

migration-reset:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}" down

migration-version:
	@go run -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest -path db/migrations -database "postgres://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?search_path=${DB_SCHEMA}" version

deploy-service:
	@sam build -t cfn/service.cfn.yaml
	@sam deploy --config-file service.sam.toml --parameter-overrides ContainerImage=269733607457.dkr.ecr.eu-central-1.amazonaws.com/users-service:latest

deploy-config:
	@sam build -t cfn/service.cfn.yaml
	@sam deploy --config-file config.sam.toml --parameter-overrides FunctionName=UsersConfigFunction OnlyConfig=1

deploy-db:
	@sam build -t cfn/db.cfn.yaml
	@sam deploy --config-file db.sam.toml

deploy-project:
	@sam build -t cfn/project.cfn.yaml
	@sam deploy --config-file project.sam.toml

lint-deploy:
	@sam validate -t cfn/service.cfn.yaml --lint
	@sam validate -t cfn/db.cfn.yaml --lint
	@sam validate -t cfn/project.cfn.yaml --lint
