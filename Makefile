# Makefile for CNF Application

.PHONY: build test run docker-build docker-push clean

# Build the application
build:
	go build -o bin/cnf-simulator main.go

# Run tests
test:
	go test -v ./...

# Run the application locally
run: build
	./bin/cnf-simulator

# Build Docker image
docker-build:
	docker build -t cnf-simulator:latest .

# Push Docker image
docker-push: docker-build
	docker tag cnf-simulator:latest $(DOCKER_REGISTRY)/cnf-simulator:$(TAG)
	docker push $(DOCKER_REGISTRY)/cnf-simulator:$(TAG)

# Clean build artifacts
clean:
	rm -rf bin/

# Install dependencies
deps:
	go mod tidy