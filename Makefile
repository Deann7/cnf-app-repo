# Makefile for CNF Application
# Defines common tasks for building, testing, and deploying the CNF application

.PHONY: build test run clean docker-build docker-push security-scan quality-check help

# Variables
BINARY_NAME=cnf-simulator
DOCKER_REGISTRY ?= ghcr.io
DOCKER_ORG ?= $(GITHUB_REPOSITORY_OWNER)
DOCKER_IMAGE ?= cnf-simulator
TAG ?= latest

# Help target
help:
	@echo "Available targets:"
	@echo "  build          - Build the application binary"
	@echo "  test           - Run unit tests"
	@echo "  run            - Run the application locally"
	@echo "  clean          - Remove build artifacts"
	@echo "  docker-build   - Build Docker image"
	@echo "  docker-push    - Push Docker image to registry"
	@echo "  security-scan  - Run security scanning"
	@echo "  quality-check  - Run quality checks"

# Build the application
build:
	@echo "Building $(BINARY_NAME)..."
	go mod tidy
	go build -o $(BINARY_NAME) main.go
	@echo "Build completed successfully!"

# Run tests
test:
	@echo "Running tests..."
	go test -v ./...
	go tool cover -func=coverage.out

# Run the application locally
run: build
	@echo "Running $(BINARY_NAME)..."
	./$(BINARY_NAME)

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(BINARY_NAME)
	rm -f coverage.out

# Build Docker image
docker-build:
	@echo "Building Docker image $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_IMAGE):$(TAG)..."
	docker build -t $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_IMAGE):$(TAG) .

# Push Docker image
docker-push: docker-build
	@echo "Pushing Docker image to registry..."
	docker push $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_IMAGE):$(TAG)

# Run security scan
security-scan:
	@echo "Running security scan with Trivy..."
	trivy image $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_IMAGE):$(TAG)

# Run quality checks
quality-check:
	@echo "Running quality checks..."
	# Check code formatting
	gofmt -l .
	# Run static analysis
	golangci-lint run
	# Run tests
	make test

# Install dependencies
deps:
	@echo "Installing dependencies..."
	go mod download

# Generate documentation
docs:
	@echo "Generating documentation..."
	godoc -http=:6060 &
	@echo "Documentation server started on http://localhost:6060"

# Run local development server
dev:
	@echo "Starting development server..."
	PORT=8080 ENVIRONMENT=development KUBERNETES_NODE_NAME=local-node go run main.go

# Build and run in Docker
docker-dev: docker-build
	@echo "Running in Docker container..."
	docker run -p 8080:8080 -e PORT=8080 -e ENVIRONMENT=development -e KUBERNETES_NODE_NAME=docker-container $(DOCKER_REGISTRY)/$(DOCKER_ORG)/$(DOCKER_IMAGE):$(TAG)

# Run all checks
check-all: deps test quality-check security-scan
	@echo "All checks passed!"

# Default target
all: build