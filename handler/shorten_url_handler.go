package handler

import (
	"encoding/json"
	"net/http"
)

// ShortenURL handles POST requests to shorten URLs
func ShortenURL(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the URL
	var requestBody struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// TODO: Implement URL shortening logic

	// Respond with a placeholder shortened URL for now
	response := struct {
		ShortenedURL string `json:"shortened_url"`
	}{
		ShortenedURL: "http://localhost:8080/abc123",
	}

	// Send the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
