.PHONY: all start build run test clean help install uninstall deps fmt lint vet docs release example-analyze example-large

# Variables
BINARY_NAME=storage-optimizer
VERSION=v1.0.0
BUILD_TIME=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GIT_COMMIT=$(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")

all: clean deps build test ## Complete build, test, and check

start: build ## Build and start the application
	@echo "Starting $(BINARY_NAME)..."
	@./$(BINARY_NAME) --help

help: ## Display help
	@echo "Storage Optimizer - Build and Development Commands"
	@echo "====================================================="
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

build: ## Build the application
	@echo "Building..."
	@go build -ldflags "-X main.Version=$(VERSION) -X main.BuildTime=$(BUILD_TIME) -X main.GitCommit=$(GIT_COMMIT)" -o $(BINARY_NAME) main.go
	@echo "Build successful: $(BINARY_NAME)"

run: build ## Build and run
	@echo "Running..."
	@./$(BINARY_NAME)

test: ## Run tests
	@echo "Running tests..."
	@go test ./...

clean: ## Clean temporary files
	@echo "Cleaning..."
	@rm -f $(BINARY_NAME)
	@go clean
	@echo "Clean complete"

fmt: ## Format code
	@echo "Formatting code..."
	@go fmt ./...

lint: ## Lint code quality
	@echo "Linting code..."
	@golangci-lint run ./... || go vet ./...

vet: ## Vet for errors
	@echo "Running vet..."
	@go vet ./...

install: build ## Install to system
	@echo "Installing..."
	@cp $(BINARY_NAME) /usr/local/bin/
	@echo "Installed: /usr/local/bin/$(BINARY_NAME)"

uninstall: ## Uninstall from system
	@echo "Uninstalling..."
	@rm -f /usr/local/bin/$(BINARY_NAME)
	@echo "Uninstalled"

deps: ## Install dependencies
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

docs: ## Generate documentation
	@echo "Generating documentation..."
	@./$(BINARY_NAME) --help

example-analyze: build ## Example: Analyze
	@./$(BINARY_NAME) analyze .

example-large: build ## Example: Large files
	@./$(BINARY_NAME) large . --limit 10

release: clean test build ## Prepare release
	@echo "Release: $(VERSION)"
	@echo "Time: $(BUILD_TIME)"
	@echo "Commit: $(GIT_COMMIT)"
