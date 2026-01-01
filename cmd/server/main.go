package main

import (
	"encoding/json"
	"fmt"
	"go-search-tool/internal/engine"
	"log"
	"net/http"
    "path/filepath"
)

// HealthResponse represents the health check response
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// SearchResponse represents the search results
type SearchResponse struct {
	Results []string `json:"results"`
	Count   int      `json:"count"`
	Error   string   `json:"error,omitempty"`
}

// enableCORS sets up the CORS headers for the response
func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// handleHealth is a simple health check handler
func handleHealth(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	response := HealthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// handleSearch processes search requests
func handleSearch(w http.ResponseWriter, r *http.Request) {
	enableCORS(&w)
	if r.Method == "OPTIONS" {
		return
	}

	path := r.URL.Query().Get("path")
	name := r.URL.Query().Get("name")
	ext := r.URL.Query().Get("ext")

    // Default path to current directory if empty
    if path == "" {
        path = "."
    }
    
    // Resolve absolute path for clarity
    absPath, err := filepath.Abs(path)
    if err == nil {
        path = absPath
    }

	results, err := engine.FindFiles(path, name, ext)
	
	response := SearchResponse{
		Results: results,
		Count:   len(results),
	}

	if err != nil {
		response.Error = err.Error()
	}
    
    // Ensure results is empty list not null
    if response.Results == nil {
        response.Results = []string{}
    }

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register routes
	http.HandleFunc("/api/health", handleHealth)
	http.HandleFunc("/api/search", handleSearch)

	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
}
