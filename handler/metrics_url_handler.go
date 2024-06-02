package handler

import (
	"encoding/json"
	"go-url-shortener-service/storage"
	"net/http"
)

// TopDomainsHandler returns a handler function that retrieves the top domains
func TopDomainsHandler(urlStore *storage.URLStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//method to retrieve the top 3 domains with the highest counts
		topDomains := urlStore.GetTopDomains(3)
		//domainCounts map to store the counts for the top 3 domains
		domainCounts := make(map[string]int)

		//Iterating over the topDomains slice and updating the domainCounts map with the counts for each domain
		for _, domain := range topDomains {
			domainCounts[domain] = urlStore.DomainCounts()[domain]
		}

		w.Header().Set("Content-Type", "application/json")
		//Encoding the domainCounts map as JSON and sending it as the response to the user
		json.NewEncoder(w).Encode(domainCounts)
	}
}
