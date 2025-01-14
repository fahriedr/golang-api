build:
	@go build -o ./tmp/main ./cmd/main.go

test:
	@go test -v ./...

run: build
	@./tmp/main

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	@go run cmd/migrate/migrations/main.go up

migrate-down:
	@go run cmd/migrate/migrations/main.go down