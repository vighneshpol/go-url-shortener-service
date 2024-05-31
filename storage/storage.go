package storage

// URLStore is a store for URL mappings we will get
type URLStore struct {
	originalToShort map[string]string // original URL -> short code
	shortToOriginal map[string]string // short code -> original URL

}

// NewURLStore creates a new URLStore from the new URL inputs
func NewURLStore() *URLStore {
	return &URLStore{
		originalToShort: make(map[string]string),
		shortToOriginal: make(map[string]string),

	}
}

// SaveMapping saves the mapping between original URL and short code
func (s *URLStore) SaveMapping(originalURL, shortCode string) {
	s.originalToShort[originalURL] = shortCode
	s.shortToOriginal[shortCode] = originalURL

}

// GetShortCode retrieves the short code for the given original URL in alphanumeric format
func (s *URLStore) GetShortCode(originalURL string) (string, bool) {
	shortCode, exists := s.originalToShort[originalURL]
	return shortCode, exists
}

// GetOriginalURL retrieves the original URL for the given short code
func (s *URLStore) GetOriginalURL(shortCode string) (string, bool) {
	originalURL, exists := s.shortToOriginal[shortCode]
	return originalURL, exists
}


