package main

import (
	"log"
	"net/http"
	"os"
	"url_shortener_server/routes"
	"url_shortener_server/shortener"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file:", err)
	}
	shortener := shortener.NewShortener()
	mux := routes.SetupRoutes(shortener)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	log.Printf("Server listening on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
