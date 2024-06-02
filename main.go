package main

import (
	"go-url-shortener-service/handler"
	"go-url-shortener-service/storage"
	"log"
	"net/http"
)

func main() {
	// Setup URL storage
	urlStore := storage.NewURLStore()

	// Define the route for shortening URLs and pass urlStore to the handler
	http.HandleFunc("/shorten", handler.ShortenURL(urlStore))

	// Define the route for redirecting shortened URLs and pass urlStore to the redirect handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		redirect(w, r, urlStore)
	})


	// Define the route for retrieving metrics
	http.HandleFunc("/metrics", handler.TopDomainsHandler(urlStore))

	// Start the HTTP server on port 8080
	log.Println("Starting URL shortener service on http://localhost:8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func redirect(w http.ResponseWriter, r *http.Request, urlStore *storage.URLStore) {
	//slice expression is used to remove the leading slash and extract the actual short code from the URL path
	shortCode := r.URL.Path[1:]

	// Look up the original URL in the store
	originalURL, exists := urlStore.GetOriginalURL(shortCode)
	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	// Redirect to the original URL
	http.Redirect(w, r, originalURL, http.StatusFound)
}
