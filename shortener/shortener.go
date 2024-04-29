package shortener

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"
	constant "url_shortener_server/constants"
)

type Shortener struct {
	shortToOriginal map[string]string
	originalToShort map[string]string
	stats           map[string]int
	mu              sync.RWMutex
	counter         int
}

func NewShortener() *Shortener {
	return &Shortener{
		shortToOriginal: make(map[string]string),
		originalToShort: make(map[string]string),
		stats:           make(map[string]int),
		counter:         0,
	}
}

const (
	base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	base        = 62
	shortLength = 6
)

func (s *Shortener) ShortenURL(w http.ResponseWriter, r *http.Request) {
	var req struct {
		URL string `json:"url"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		log.Println("Error decoding JSON request:", err)
		return
	}

	// Validate URL
	if req.URL == "" {
		http.Error(w, constant.EmptyShortURL, http.StatusBadRequest)
		log.Println("Received empty URL in ShortenURL")
		return
	}

	// Check if URL already exists
	s.mu.RLock()
	shortURL, exists := s.originalToShort[req.URL]
	s.mu.RUnlock()
	if exists {
		resp := struct {
			ShortURL string `json:"short_url"`
		}{
			ShortURL: shortURL,
		}
		json.NewEncoder(w).Encode(resp)
		log.Printf("Short URL already exists for %s: %s\n", req.URL, shortURL)
		return
	}

	// Generate short URL
	shortURL = s.generateShortURL(req.URL)

	s.mu.Lock()
	s.originalToShort[req.URL] = shortURL
	s.mu.Unlock()

	log.Printf("Short URL generated for %s: %s\n", req.URL, shortURL)

	resp := struct {
		ShortURL string `json:"short_url"`
	}{
		ShortURL: shortURL,
	}
	json.NewEncoder(w).Encode(resp)
}

func (s *Shortener) generateShortURL(originalURL string) string {
	shortCode := s.toBase62(s.counter)
	s.counter++

	s.mu.Lock()
	defer s.mu.Unlock()
	s.shortToOriginal[shortCode] = originalURL
	return "/short/" + shortCode
}

func (s *Shortener) toBase62(num int) string {
	var sb strings.Builder
	for num > 0 {
		remainder := num % base
		sb.WriteByte(base62Chars[remainder])
		num /= base
	}
	shortCode := sb.String()
	for i := len(shortCode); i < shortLength; i++ {
		sb.WriteByte(base62Chars[0])
	}
	return sb.String()
}
