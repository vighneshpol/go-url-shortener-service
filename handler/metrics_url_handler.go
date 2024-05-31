package handler

import (
	"go-url-shortener-service/storage"
	"net/http"
	"encoding/json"
)

// TopDomainsHandler returns a handler function that retrieves the top domains
func TopDomainsHandler(urlStore *storage.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		topDomains := urlStore.GetTopDomains(3)
		domainCounts := make(map[string]int)

		for _, domain := range topDomains {
			domainCounts[domain] = urlStore.DomainCounts()[domain]
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(domainCounts)
	}
}
