.PHONY: run
run:
	go run ./cmd/dbservice/.

.PHONY: wire
wire:
	wire ./internal/app/

.PHONY: migrate-new
migrate-new:
	migrate create -ext sql -dir ./migrations ${NAME}

.PHONY: migrate
migrate:
	migrate -path ./migrations -database "postgres://rightdesire:secret@postgres:5432/main?sslmode=disable" up 

.PHONY: migrate-down
migrate-down:
	migrate -path ./migrations -database "postgres://rightdesire:secret@postgres:5432/main?sslmode=disable" down  

.PHONY: database-drop
database-drop:
	migrate -database "postgres://rightdesire:secret@postgres:5432/main?sslmode=disable" -path ./migrations drop -f

.PHONY: lint
lint:
	golangci-lint run --out-format=colored-line-number
