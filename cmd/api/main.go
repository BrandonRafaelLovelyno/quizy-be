package main

import (
	"log"
	"net/http"
	"os"

	"quizy-be/internal/db"
	"quizy-be/internal/middleware"
	"quizy-be/internal/router"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found - using system environment variables")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable not set")
	}

	if err := db.InitMongo(mongoURI); err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}
	defer db.DisconnectMongo()

	r := chi.NewRouter()
	middleware.SetupMiddleware(r)
	router.SetupRouter(r)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
