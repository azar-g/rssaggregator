package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/azar-g/rssaggregator/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	// Load environment variables from .env file
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL environment variable is not set")
	}
	fmt.Printf("Server is running on port: %s\n", port)

	// Initialize database connection
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}

	// apiCfg holds the API configuration settings for the application.
	// It initializes an instance of apiConfig by setting up the database component,
	// where the DB field is assigned a new database connection obtained via database.New.
	// This configuration is critical for enabling API endpoints to interact with the underlying database.
	apiCfg := apiConfig{
		DB: database.New(db),
	}

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
	v1Router.Post("/users", apiCfg.createUserHandler)
	v1Router.Get("/users", apiCfg.authMiddleware(apiCfg.getUser))
	v1Router.Post("/feeds", apiCfg.authMiddleware(apiCfg.createUserFeed))
	v1Router.Get("/feeds", apiCfg.getAllFeeds)
	v1Router.Post("/feed-follows", apiCfg.authMiddleware(apiCfg.createUserFeedFollow))
	v1Router.Get("/feed-follows", apiCfg.authMiddleware(apiCfg.getAllFeedFollowsByUser))
	v1Router.Delete("/feed-follows/{id}", apiCfg.authMiddleware(apiCfg.deleteFeedFollow))

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router,
	}
	log.Printf("Starting server on port %v", port)
	// Start the server
	err = server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

	router.Get("/go", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the RSS Aggregator!"))
	})
}
