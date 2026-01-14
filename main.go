package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// CNFStatus represents the status of our Cloud-Native Network Function
type CNFStatus struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Status      string    `json:"status"`
	StartedAt   time.Time `json:"started_at"`
	Environment string    `json:"environment"`
	K8sNode     string    `json:"k8s_node"`
}

// Global variable to store CNF status
var cnfStatus CNFStatus

// initializes the CNF status with default values
func init() {
nodeName := os.Getenv("KUBERNETES_NODE_NAME")
if nodeName == "" {
	nodeName = "unknown-node"
}

cnfStatus = CNFStatus{
	ID:          generateID(),
	Name:        "Simple-CNFSimulator",
	Version:     "1.0.0",
	Status:      "running",
	StartedAt:   time.Now(),
	Environment: os.Getenv("ENVIRONMENT"),
	K8sNode:     nodeName,
}
}

// generateID creates a simple unique ID for the CNF instance
func generateID() string {
return fmt.Sprintf("cnf-%d", time.Now().Unix())
}

// healthHandler returns the health status of the CNF
func healthHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")
fmt.Fprintf(w, `{"status": "healthy", "service": "cnf-simulator", "timestamp": "%s"}`, time.Now().Format(time.RFC3339))
}

// statusHandler returns detailed status information about the CNF
func statusHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

// Update status with current timestamp
cnfStatus.Status = "running"

fmt.Fprintf(w, `{
	"id": "%s",
	"name": "%s",
	"version": "%s",
	"status": "%s",
	"started_at": "%s",
	"environment": "%s",
	"k8s_node": "%s",
	"current_time": "%s"
}`,
cnfStatus.ID,
cnfStatus.Name,
cnfStatus.Version,
cnfStatus.Status,
cnfStatus.StartedAt.Format(time.RFC3339),
cnfStatus.Environment,
cnfStatus.K8sNode,
time.Now().Format(time.RFC3339))
}

// configHandler displays environment configuration
func configHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

envVarsJSON := "{"
for _, env := range os.Environ() {
	parts := strings.SplitN(env, "=", 2)
	if len(parts) == 2 {
		if envVarsJSON != "{" {
			envVarsJSON += ","
		}
		envVarsJSON += fmt.Sprintf("\n    \"%s\": \"%s\"", parts[0], parts[1])
	}
}
envVarsJSON += "\n  }"

fmt.Fprintf(w, `{
	"config": {
		"port": "%s",
		"environment": "%s",
		"kubernetes_node": "%s"
	},
	"env_vars": %s
}`,
os.Getenv("PORT"),
os.Getenv("ENVIRONMENT"),
os.Getenv("KUBERNETES_NODE_NAME"),
envVarsJSON)
}

// infoHandler provides general information about the CNF
func infoHandler(w http.ResponseWriter, r *http.Request) {
w.Header().Set("Content-Type", "application/json")

fmt.Fprintf(w, `{
	"service": "Cloud-Native Network Function Simulator",
	"description": "A simple Go application simulating a CNF for O-Cloud environment",
	"endpoints": [
		"/health - Health check endpoint",
		"/status - Detailed status information",
		"/config - Configuration information",
		"/info - Service information"
	],
	"version": "1.0.0",
	"author": "O-Cloud CNF Simulator"
}`)
}

func main() {
port := os.Getenv("PORT")
if port == "" {
	port = "8080"
}

fmt.Printf("Starting CNF Simulator on port %s\n", port)
fmt.Printf("CNF Instance ID: %s\n", cnfStatus.ID)
fmt.Printf("Running on Kubernetes Node: %s\n", cnfStatus.K8sNode)
fmt.Printf("Environment: %s\n", cnfStatus.Environment)

// Define HTTP routes
http.HandleFunc("/health", healthHandler)
http.HandleFunc("/status", statusHandler)
http.HandleFunc("/config", configHandler)
http.HandleFunc("/info", infoHandler)

// Default handler
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	statusHandler(w, r)
})

// Start the HTTP server
log.Fatal(http.ListenAndServe(":"+port, nil))
}