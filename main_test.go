package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRootHandler(t *testing.T) {
	mux := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	if !strings.Contains(rr.Body.String(), "Go Web Application") {
		t.Errorf("Expected body to contain 'Go Web Application'")
	}
}

func TestHealthHandler(t *testing.T) {
	mux := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/healthz", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var resp HealthResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to parse health response: %v", err)
	}

	if resp.Status != "UP" {
		t.Errorf("Expected status 'UP', got %s", resp.Status)
	}
}

func TestInfoHandler(t *testing.T) {
	mux := setupRouter()

	req, err := http.NewRequest(http.MethodGet, "/api/info", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, rr.Code)
	}

	var resp InfoResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("Failed to parse info response: %v", err)
	}

	if resp.AppName != "Go Multi-Branch Demo App" {
		t.Errorf("Expected AppName 'Go Multi-Branch Demo App', got %s", resp.AppName)
	}
}
