package handler

import (

	"go-url-shortener-service/storage"
	"net/http"
)

// TopDomainsHandler returns a handler function that retrieves the top domains
func TopDomainsHandler(urlStore *storage.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		
	}
}
