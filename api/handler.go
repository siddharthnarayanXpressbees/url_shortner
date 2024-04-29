package api

import (
	"net/http"
	"url_shortener_server/shortener"
)

func ShortenHandler(s *shortener.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.ShortenURL(w, r)
	}
}

func RedirectHandler(s *shortener.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Redirect(w, r)
	}
}

func MetricsHandler(s *shortener.Shortener) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Metrics(w, r)
	}
}
