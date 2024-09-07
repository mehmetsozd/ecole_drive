APP_NAME = ecole_drive
MAIN_SRC = cmd/main.go
MIGRATE_SRC = cmd/migrate/main.go

all: run

run:
	@echo "==> Running the application..."
	@go run $(MAIN_SRC)

build:
	@echo "==> Building the project..."
	go build -o $(APP_NAME) $(MAIN_SRC)

migrate-up:
	@echo "==> Running migrations up..."
	@go run $(MIGRATE_SRC) up

migrate-down:
	@echo "==> Running migrations down..."
	@go run $(MIGRATE_SRC) down

deps:
	@echo "==> Installing dependencies..."
	go mod tidy

clean:
	@echo "==> Cleaning up..."
	@if [ -f $(APP_NAME) ]; then \
		rm $(APP_NAME); \
		echo "==> Binary file removed."; \
	else \
		echo "==> No binary file found."; \
	fi

test:
	@echo "==> Running tests..."
	go test ./...

help:
	@echo "==> Available targets:"
	@echo "    run            - Run the application"
	@echo "    build          - Build the project"
	@echo "    migrate-up     - Apply database migrations (up)"
	@echo "    migrate-down   - Roll back database migrations (down)"
	@echo "    deps           - Install dependencies"
	@echo "    clean          - Clean up build artifacts"
	@echo "    test           - Run all tests"
