package handler

import (
	"encoding/json"
	"net/http"
	"go-url-shortener-service/storage"
	"go-url-shortener-service/shortener"
)

// ShortenURL returns a handler function that shortens URLs
func ShortenURL(urlStore *storage.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body to get the URL
		var requestBody struct {
			URL string `json:"url"`
		}
		if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		// Check if the URL has already been shortened
		if shortCode, exists := urlStore.GetShortCode(requestBody.URL); exists {
			// Respond with the existing shortened URL
			shortenedURL := "http://localhost:8080/" + shortCode
			response := struct {
				ShortenedURL string `json:"shortened_url"`
			}{
				ShortenedURL: shortenedURL,
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
			return
		}

		// Generate short code for the URL
		shortCode := shortener.GenerateShortCode(6)

		// Save the mapping in the store
		urlStore.SaveMapping(requestBody.URL, shortCode)

		// Construct the shortened URL
		shortenedURL := "http://localhost:8080/" + shortCode

		// Respond with the shortened URL
		response := struct {
			ShortenedURL string `json:"shortened_url"`
		}{
			ShortenedURL: shortenedURL,
		}

		// Send the response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
