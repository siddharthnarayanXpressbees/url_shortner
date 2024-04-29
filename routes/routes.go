package routes

import (
	"log"
	"net/http"
	"url_shortener_server/api"
	"url_shortener_server/shortener"
)

func SetupRoutes(s *shortener.Shortener) *http.ServeMux {
	mux := http.NewServeMux()
	addRoute(mux, http.MethodPost, "/shorten", api.ShortenHandler(s))
	addRoute(mux, http.MethodGet, "/short/", api.RedirectHandler(s))
	addRoute(mux, http.MethodGet, "/metrics", api.MetricsHandler(s))

	return mux
}

func addRoute(mux *http.ServeMux, method, path string, handler http.HandlerFunc) {
	mux.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			handler(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			log.Println("Method Not Allowed")
		}
	})
}
