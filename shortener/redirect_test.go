package shortener

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRedirect(t *testing.T) {
	shortener := NewShortener()
	shortURL := "http://localhost:8080/short/aaaa"
	shortcode := "aaaa"

	// Add a test short URL to the shortener
	shortener.shortToOriginal[shortcode] = "https://www.google.com"

	req, err := http.NewRequest("GET", shortURL, nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.Redirect)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusFound)
	}
}
