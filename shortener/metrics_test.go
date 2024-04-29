package shortener

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

var responseStats []struct {
	Domain string `json:"Domain"`
	Count  int    `json:"Count"`
}

func TestMetrics(t *testing.T) {
	shortener := NewShortener()

	// Add test data to the shortener's stats map
	shortener.stats["www.flipkart.com"] = 5
	shortener.stats["www.google.com"] = 3
	shortener.stats["www.amazon.com"] = 7

	req, err := http.NewRequest("GET", "/metrics", nil)
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler := http.HandlerFunc(shortener.Metrics)
	handler.ServeHTTP(response, req)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	if err := json.Unmarshal(response.Body.Bytes(), &responseStats); err != nil {
		t.Fatal(err)
	}

	responseStatsMap := make(map[string]int)
	for _, stats := range responseStats {
		responseStatsMap[stats.Domain] = stats.Count
	}

	if responseStatsMap["www.flipkart.com"] != 5 && responseStatsMap["www.google.com"] != 3 && responseStatsMap["www.amazon.com"] != 7 {
		t.Errorf("handler returned unexpected value miss-match: ")
	}
}
