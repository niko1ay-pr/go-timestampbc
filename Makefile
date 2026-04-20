.PHONY: run setup test build clean migrate docker-build

BINARY_NAME=timestampbc
BUILD_DIR=./bin

setup:
	@echo "Setting up environment..."
	@if [ ! -f .env ]; then \
		cp .env.example .env; \
		echo ".env file created from .env.example"; \
		echo "Edit .env file to configure your environment"; \
	else \
		echo ".env file already exists"; \
	fi
	@mkdir -p ./data
	@echo "Setup complete!"

run:
	@echo "Starting server..."
	go run ./cmd/api

dev:
	@if command -v air > /dev/null; then \
        air -c .air.toml; \
    else \
        echo "Air not installed. Installing..."; \
        go install github.com/air-verse/air@v1.49.0; \
        air -c .air.toml; \
    fi

test:
	go test ./... -v

build:
	@mkdir -p $(BUILD_DIR)
	go build -o $(BUILD_DIR)/$(BINARY_NAME) ./cmd/api

clean:
	rm -rf $(BUILD_DIR)
	rm -rf coverage.out
	rm -rf ./data/*.db

migrate:
	@echo "Migrations will be here soon"

tidy:
	go mod tidy
