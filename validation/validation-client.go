// validation-client.go - Go client for comprehensive API validation
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HealthResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Version   string    `json:"version"`
}

type StatusResponse struct {
	AppName       string            `json:"app_name"`
	InstanceID    string            `json:"instance_id"`
	StartTime     time.Time         `json:"start_time"`
	UptimeSeconds int               `json:"uptime_seconds"`
	Config        map[string]string `json:"config"`
}

func validateHealthEndpoint(baseURL string) error {
	resp, err := http.Get(fmt.Sprintf("%s/health", baseURL))
	if err != nil {
		return fmt.Errorf("failed to connect to health endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("health endpoint returned status code %d", resp.StatusCode)
	}

	var health HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&health); err != nil {
		return fmt.Errorf("failed to parse health response: %v", err)
	}

	if health.Status != "healthy" {
		return fmt.Errorf("health status is not 'healthy': got '%s'", health.Status)
	}

	fmt.Printf("✓ Health endpoint validation passed: %s\n", health.Status)
	return nil
}

func validateStatusEndpoint(baseURL string) error {
	resp, err := http.Get(fmt.Sprintf("%s/status", baseURL))
	if err != nil {
		return fmt.Errorf("failed to connect to status endpoint: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("status endpoint returned status code %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read status response: %v", err)
	}

	var status StatusResponse
	if err := json.Unmarshal(body, &status); err != nil {
		return fmt.Errorf("failed to parse status response: %v", err)
	}

	if status.AppName == "" {
		return fmt.Errorf("status response missing app_name field")
	}

	fmt.Printf("✓ Status endpoint validation passed: %s\n", status.AppName)
	return nil
}

func main() {
	baseURL := "http://localhost:8080" // Should be replaced with actual service URL

	fmt.Println("Starting API endpoint validation...")

	if err := validateHealthEndpoint(baseURL); err != nil {
		fmt.Printf("✗ Health validation failed: %v\n", err)
		return
	}

	if err := validateStatusEndpoint(baseURL); err != nil {
		fmt.Printf("✗ Status validation failed: %v\n", err)
		return
	}

	fmt.Println("All API endpoint validations passed!")
}
