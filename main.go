package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from .env file
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	fmt.Printf("Server is running on port: %s\n", port)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum age in seconds for preflight requests
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", readinessHandler)
	v1Router.Get("/error", errorHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	log.Printf("Starting server on port %v", port)
	// Start the server
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

	router.Get("/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the RSS Aggregator!"))
	})
}