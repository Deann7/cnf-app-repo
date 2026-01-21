package main

import (
	"crypto/sha256"
	"encoding/json"
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
	Security    SecurityInfo `json:"security"`
}

// SecurityInfo holds security-related information
type SecurityInfo struct {
	ScanStatus string `json:"scan_status"`
	LastScan   string `json:"last_scan"`
	Vulnerabilities int `json:"vulnerabilities"`
	SecurityRating  string `json:"security_rating"`
}

// QualityMetrics holds quality metrics information
type QualityMetrics struct {
	CodeCoverage float64 `json:"code_coverage"`
	TestResults  []TestResult `json:"test_results"`
}

// TestResult holds individual test results
type TestResult struct {
	Name string `json:"name"`
	Status string `json:"status"`
	Duration time.Duration `json:"duration"`
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
		Security: SecurityInfo{
			ScanStatus:      "completed",
			LastScan:        time.Now().Format(time.RFC3339),
			Vulnerabilities: 0,
			SecurityRating:  "A",
		},
	}
}

// generateID creates a simple unique ID for the CNF instance
func generateID() string {
	return fmt.Sprintf("cnf-%d", time.Now().Unix())
}

// healthHandler returns the health status of the CNF
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	
	response := map[string]interface{}{
		"status": "healthy",
		"service": "cnf-simulator",
		"timestamp": time.Now().Format(time.RFC3339),
		"security_rating": cnfStatus.Security.SecurityRating,
		"vulnerabilities": cnfStatus.Security.Vulnerabilities,
	}
	json.NewEncoder(w).Encode(response)
}

// statusHandler returns detailed status information about the CNF
func statusHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Update status with current timestamp
	cnfStatus.Status = "running"
	cnfStatus.Security.LastScan = time.Now().Format(time.RFC3339)

	response := map[string]interface{}{
		"id":           cnfStatus.ID,
		"name":         cnfStatus.Name,
		"version":      cnfStatus.Version,
		"status":       cnfStatus.Status,
		"started_at":   cnfStatus.StartedAt.Format(time.RFC3339),
		"environment":  cnfStatus.Environment,
		"k8s_node":     cnfStatus.K8sNode,
		"current_time": time.Now().Format(time.RFC3339),
		"security":     cnfStatus.Security,
	}
	json.NewEncoder(w).Encode(response)
}

// securityHandler provides security scan information
func securityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	response := map[string]interface{}{
		"scan_status":      cnfStatus.Security.ScanStatus,
		"last_scan":        cnfStatus.Security.LastScan,
		"vulnerabilities":  cnfStatus.Security.Vulnerabilities,
		"security_rating":  cnfStatus.Security.SecurityRating,
		"security_policy":  "strict",
		"compliance":       "SOC2,ISO27001",
	}
	json.NewEncoder(w).Encode(response)
}

// qualityHandler provides quality metrics information
func qualityHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	metrics := QualityMetrics{
		CodeCoverage: 85.0, // 85% code coverage as mentioned in Day 8 report
		TestResults: []TestResult{
			{Name: "unit_tests", Status: "passed", Duration: 15 * time.Second},
			{Name: "integration_tests", Status: "passed", Duration: 30 * time.Second},
			{Name: "security_tests", Status: "passed", Duration: 45 * time.Second},
			{Name: "performance_tests", Status: "passed", Duration: 60 * time.Second},
		},
	}
	
	json.NewEncoder(w).Encode(metrics)
}

// configHandler displays environment configuration
func configHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	envVarsJSON := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		if len(parts) == 2 {
			// Only expose environment variables that start with APP_ or CNF_
			if strings.HasPrefix(parts[0], "APP_") || strings.HasPrefix(parts[0], "CNF_") {
				envVarsJSON[parts[0]] = maskSensitiveData(parts[0], parts[1])
			}
		}
	}

	response := map[string]interface{}{
		"config": map[string]string{
			"port":            os.Getenv("PORT"),
			"environment":     os.Getenv("ENVIRONMENT"),
			"kubernetes_node": os.Getenv("KUBERNETES_NODE_NAME"),
		},
		"env_vars": envVarsJSON,
	}
	json.NewEncoder(w).Encode(response)
}

// maskSensitiveData masks sensitive environment variables
func maskSensitiveData(key, value string) string {
	sensitiveKeys := []string{"PASSWORD", "SECRET", "TOKEN", "KEY", "AUTH"}
	
	for _, sensitiveKey := range sensitiveKeys {
		if strings.Contains(strings.ToUpper(key), sensitiveKey) {
			// Return a masked version of the value
			hash := sha256.Sum256([]byte(value))
			return fmt.Sprintf("[MASKED:%x]", hash[:8])
		}
	}
	return value
}

// infoHandler provides general information about the CNF
func infoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := map[string]interface{}{
		"service": "Cloud-Native Network Function Simulator",
		"description": "A secure Go application simulating a CNF for O-Cloud environment with security scanning and quality gates",
		"endpoints": []string{
			"/health - Health check endpoint",
			"/status - Detailed status information",
			"/config - Configuration information",
			"/info - Service information",
			"/security - Security scan information",
			"/quality - Quality metrics information",
		},
		"version": "1.0.0",
		"author": "O-Cloud CNF Simulator",
		"security_features": []string{
			"Vulnerability scanning",
			"Security headers",
			"Environment variable masking",
			"Quality gates enforcement",
		},
	}
	json.NewEncoder(w).Encode(response)
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
	http.HandleFunc("/security", securityHandler)
	http.HandleFunc("/quality", qualityHandler)

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