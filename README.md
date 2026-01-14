# Cloud-Native Network Function (CNF) Simulator

This is a simple Go application that simulates a Cloud-Native Network Function (CNF) for the O-Cloud environment. It provides various endpoints to monitor and manage the simulated CNF.

## Features

- Health check endpoint (`/health`)
- Status information endpoint (`/status`)
- Configuration information endpoint (`/config`)
- Service information endpoint (`/info`)
- Environment variable monitoring
- Kubernetes node identification

## Endpoints

- `/health` - Returns health status of the CNF
- `/status` - Provides detailed status information
- `/config` - Shows environment configuration
- `/info` - Displays service information
- `/` - Root endpoint that returns status information

## Environment Variables

- `PORT` - Port number to run the service on (default: 8080)
- `ENVIRONMENT` - Environment name (e.g., development, staging, production)
- `KUBERNETES_NODE_NAME` - Name of the Kubernetes node where the CNF is running

## Building and Running

To build and run the application:

```bash
go mod init cnf-app-repo
go mod tidy
go run main.go
```

Or build a binary:

```bash
go build -o cnf-simulator main.go
./cnf-simulator
```

## Running in Container

To containerize this application, create a Dockerfile:

```dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
RUN go build -o cnf-simulator main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/cnf-simulator .
EXPOSE 8080
CMD ["./cnf-simulator"]
```

## Purpose

This simulator is designed to demonstrate how a Cloud-Native Network Function would behave in an O-Cloud environment, providing insights into:

- Application lifecycle management
- Health monitoring
- Configuration handling
- Kubernetes integration
- Environment awareness