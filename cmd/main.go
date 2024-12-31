package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/turplespace/hlomailapi/internal/handlers"
	"github.com/turplespace/hlomailapi/internal/middleware"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Create a new ServeMux
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/api/contact", handlers.HandlePostContactFormHandler)

	// Wrap the mux with the CorsMiddleware
	log.Println("Server started on port", port)
	if err := http.ListenAndServe(":"+port, middleware.CorsMiddleware(mux)); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
