# Multi-stage build for CNF application
# First stage: Build the Go application
FROM golang:1.21-alpine AS builder

# Set recommended environment for Go builds
ENV CGO_ENABLED=0
ENV GOOS=linux

# Install git for go modules that might need it
RUN apk add --no-cache git

# Set the working directory inside the container
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies with verification
RUN go mod download

# Copy the source code
COPY main.go .

# Build the application with security and optimization flags
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o cnf-simulator .

# Second stage: Create a minimal runtime image
FROM alpine:latest

# Install only necessary packages for security
RUN apk --no-cache add ca-certificates tzdata

# Create a non-root user for security with specific UID/GID
RUN addgroup -g 65532 appgroup &&\
    adduser -D -s /bin/sh -u 65532 -G appgroup appuser

# Set the working directory
WORKDIR /app

# Copy the binary from the first stage
COPY --from=builder --chown=appuser:appgroup /app/cnf-simulator .

# Switch to non-root user
USER appuser

# Expose the port the application runs on
EXPOSE 8080

# Health check for container orchestration systems
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
  CMD wget -q -O /dev/null http://localhost:8080/health || exit 1

# Command to run the application
CMD ["./cnf-simulator"]