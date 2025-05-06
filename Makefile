 .PHONY: build run migrate-up migrate-down test docker-build docker-run

build:
	go build -o bin/person-enrichment cmd/main.go

run: build
	./bin/person-enrichment

migrate-up:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" down



docker-build:
	docker build -t person-enrichment .

docker-run:
	 swag init -g cmd/main.go --output docs/ && docker-compose down -v && docker-compose up --build