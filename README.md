# Cloud-Native Network Function (CNF) Simulator

This is a secure Go application that simulates a Cloud-Native Network Function (CNF) for the O-Cloud environment. It provides various endpoints to monitor and manage the simulated CNF with integrated security scanning and quality gates.

## Features

- Health check endpoint (`/health`)
- Status information endpoint (`/status`)
- Configuration information endpoint (`/config`)
- Service information endpoint (`/info`)
- Security scan information endpoint (`/security`)
- Quality metrics endpoint (`/quality`)
- Environment variable monitoring
- Kubernetes node identification
- Security headers and protections
- Environment variable masking

## Endpoints

- `/health` - Returns health status of the CNF
- `/status` - Provides detailed status information
- `/config` - Shows environment configuration
- `/info` - Displays service information
- `/security` - Provides security scan information
- `/quality` - Provides quality metrics information
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

## Security Features

This application implements various security measures:

- Security headers (X-Content-Type-Options, X-Frame-Options, X-XSS-Protection)
- Environment variable masking for sensitive data
- Non-root container execution
- Security scanning integration
- Vulnerability reporting

## Quality Gates

This application enforces quality standards:

- Code coverage threshold (85% minimum)
- Security rating requirements (A-grade minimum)
- Vulnerability count limits (0 critical vulnerabilities)
- Performance thresholds (response time, throughput)

## Purpose

This simulator is designed to demonstrate how a Cloud-Native Network Function would behave in an O-Cloud environment, providing insights into:

- Application lifecycle management
- Health monitoring
- Configuration handling
- Kubernetes integration
- Environment awareness
- Security scanning integration
- Quality gates enforcement

## CI/CD Integration

This application is set up with a comprehensive CI pipeline using GitHub Actions that:
- Builds the application on every push/PR
- Runs tests to ensure code quality
- Performs security scanning using Trivy, Snyk, and SonarQube
- Enforces quality gates with code coverage checks
- Builds and pushes Docker images to GitHub Container Registry
- Tags images with commit SHA for traceability
- Implements security scanning and quality gates

## Additional Features

This application now includes:
- Comprehensive security scanning capabilities
- Quality metrics tracking
- Monitoring and alerting configurations
- Container security policies
- O-Cloud platform integration
- Documentation for security implementation
- Makefile for simplified development tasks
