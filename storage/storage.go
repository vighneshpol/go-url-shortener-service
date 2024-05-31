package storage

import (
	"net/url"
	"sort"
)

// URLStore is a store for URL mappings
type URLStore struct {
	originalToShort map[string]string // original URL to short code
	shortToOriginal map[string]string // short code to original URL
	domainCounts    map[string]int    // Validation of domain count
}

// NewURLStore creates a new URLStore from the new URL inputs
func NewURLStore() *URLStore {
	return &URLStore{
		originalToShort: make(map[string]string),
		shortToOriginal: make(map[string]string),
		domainCounts:    make(map[string]int),
	}
}

// SaveMapping saves the mapping between original URL and short code
func (s *URLStore) SaveMapping(originalURL, shortCode string) {
	s.originalToShort[originalURL] = shortCode
	s.shortToOriginal[shortCode] = originalURL

	// Update domain count
	domain := getDomainFromURL(originalURL)
	s.domainCounts[domain]++
}

// GetShortCode retrieves the short code for the given original URL
func (s *URLStore) GetShortCode(originalURL string) (string, bool) {
	shortCode, exists := s.originalToShort[originalURL]
	return shortCode, exists
}

// GetOriginalURL retrieves the original URL for the given short code
func (s *URLStore) GetOriginalURL(shortCode string) (string, bool) {
	originalURL, exists := s.shortToOriginal[shortCode]
	return originalURL, exists
}

// GetTopDomains retrieves the top N domains with the highest counts
func (s *URLStore) GetTopDomains(n int) []string {
	type domainCount struct {
		Domain string
		Count  int
	}
	var counts []domainCount
	for domain, count := range s.domainCounts {
		counts = append(counts, domainCount{Domain: domain, Count: count})
	}

	// Sort by count in descending order
	sort.Slice(counts, func(i, j int) bool {
		return counts[i].Count > counts[j].Count
	})

	// Get top N domains
	var topDomains []string
	for i := 0; i < n && i < len(counts); i++ {
		topDomains = append(topDomains, counts[i].Domain)
	}
	return topDomains
}

// DomainCounts returns the domain counts map
func (s *URLStore) DomainCounts() map[string]int {
	return s.domainCounts
}

// getDomainFromURL extracts the domain from a URL string
func getDomainFromURL(rawURL string) string {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return ""
	}
	return parsedURL.Host
}
