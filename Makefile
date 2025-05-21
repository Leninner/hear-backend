.PHONY: build build-local build-prod run-prod run clean

up-db:
	docker compose up -d db

build-local:
	go build -o tmp/main cmd/main.go

build-prod:
	docker compose build

run-local:
	go run cmd/main.go

run-prod:
	docker compose up

clean:
	rm -f tmp/main
	docker compose down

