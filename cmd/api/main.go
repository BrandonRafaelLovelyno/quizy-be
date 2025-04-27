package main

import (
	"log"
	"net/http"

	"quizy-be/internal/db"
	"quizy-be/internal/middleware"
	"quizy-be/internal/router"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	cleanup, err := db.InitializeMongoDB()
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	r := chi.NewRouter()
	middleware.SetupMiddleware(r)
	router.SetupRouter(r)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
