package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
	"time"
)

// HealthServer provides HTTP health check endpoints for Kubernetes probes.
type HealthServer struct {
	server *http.Server
	ready  atomic.Bool
}

// HealthStatus represents the response for health endpoints.
type HealthStatus struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

// NewHealthServer creates a new health server on the specified port.
func NewHealthServer(port int) *HealthServer {
	h := &HealthServer{}
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.handleHealth)
	mux.HandleFunc("/ready", h.handleReady)

	h.server = &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	return h
}

// Start starts the health server in a goroutine.
func (h *HealthServer) Start() {
	go func() {
		if err := h.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Health server error: %v\n", err)
		}
	}()
}

// Stop gracefully shuts down the health server.
func (h *HealthServer) Stop(ctx context.Context) error {
	return h.server.Shutdown(ctx)
}

// SetReady marks the service as ready to receive traffic.
func (h *HealthServer) SetReady(ready bool) {
	h.ready.Store(ready)
}

// handleHealth handles liveness probe requests.
// Returns 200 OK if the service is alive.
func (h *HealthServer) handleHealth(w http.ResponseWriter, r *http.Request) {
	status := HealthStatus{
		Status:    "healthy",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(status)
}

// handleReady handles readiness probe requests.
// Returns 200 OK if the service is ready to receive traffic.
func (h *HealthServer) handleReady(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if h.ready.Load() {
		status := HealthStatus{
			Status:    "ready",
			Timestamp: time.Now().UTC().Format(time.RFC3339),
		}
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(status)
		return
	}

	status := HealthStatus{
		Status:    "not_ready",
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	w.WriteHeader(http.StatusServiceUnavailable)
	_ = json.NewEncoder(w).Encode(status)
}

// RunHealthCheck performs a health check against the local health endpoint.
// Used by the Dockerfile HEALTHCHECK command.
func RunHealthCheck(port int) error {
	url := fmt.Sprintf("http://localhost:%d/health", port)
	client := &http.Client{Timeout: 5 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("health check failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("health check returned status %d", resp.StatusCode)
	}

	return nil
}
