.PHONY: build build-local build-prod run-prod run clean dev test test-coverage migrate-up migrate-down sqlc-generate db-reset up-db install-tools fmt lint deps up down logs

# Development
dev:
	air

# Build
build-local:
	go build -o tmp/main cmd/main.go

build-prod:
	docker compose build

# Run
run-local:
	go run cmd/main.go

run-prod:
	docker compose up -d

# Docker Compose Management
up:
	docker compose up --build -d

down:
	docker compose down

logs:
	docker compose logs -f

# Database
up-db:
	docker compose up -d db

migrate-up:
	migrate -path db/migrations -database "postgres://hear:hear@localhost:5432/hear?sslmode=disable" up

migrate-down:
	migrate -path db/migrations -database "postgres://hear:hear@localhost:5432/hear?sslmode=disable" down

sqlc-generate:
	sqlc generate

db-reset:
	migrate -path db/migrations -database "postgres://hear:hear@localhost:5432/hear?sslmode=disable" down
	migrate -path db/migrations -database "postgres://hear:hear@localhost:5432/hear?sslmode=disable" up

# Testing
test:
	go test ./...

test-coverage:
	go test -cover ./...

# Cleanup
clean:
	rm -f tmp/main
	docker compose down

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest
	go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

# Format and lint
fmt:
	gofmt -s -w .

lint:
	golangci-lint run

# Dependencies
deps:
	go mod download
	go mod tidy

