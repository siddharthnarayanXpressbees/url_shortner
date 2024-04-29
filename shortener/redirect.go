package shortener

import (
	"net/http"
	"strings"
	constant "url_shortener_server/constants"
)

func (s *Shortener) Redirect(w http.ResponseWriter, r *http.Request) {
	shortCode := r.URL.Path[len(constant.ShortURLPrefix):]
	if shortCode == "" {
		http.NotFound(w, r)
		return
	}

	s.mu.RLock()
	originalURL, exists := s.shortToOriginal[shortCode]
	s.mu.RUnlock()
	if !exists {
		http.NotFound(w, r)
		return
	}

	domain := extractDomain(originalURL)
	s.mu.Lock()
	s.stats[domain]++
	s.mu.Unlock()

	http.Redirect(w, r, originalURL, http.StatusFound)
}

func extractDomain(url string) string {
	start := strings.Index(url, "://")
	if start == -1 {
		return ""
	}
	end := strings.Index(url[start+3:], "/")
	if end == -1 {
		return url[start+3:]
	}
	return url[start+3 : start+3+end]
}
