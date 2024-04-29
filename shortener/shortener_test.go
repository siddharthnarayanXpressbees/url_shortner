// shortener/shortener_test.go
package shortener

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShortenURL(t *testing.T) {
	shortener := NewShortener()

	req, err := http.NewRequest("POST", "/shorten", strings.NewReader(`{"url": "https://www.google.com"}`))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.ShortenURL)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if the response body contains a valid short URL
	expectedPrefix := `{"short_url":"/short/`
	if !strings.HasPrefix(response.Body.String(), expectedPrefix) {
		t.Errorf("handler returned unexpected body: got %v want prefix %v",
			response.Body.String(), expectedPrefix)
	}
}
