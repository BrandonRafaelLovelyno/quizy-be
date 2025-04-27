package router

import (
	"quizy-be/internal/db"
	"quizy-be/internal/handlers"
	"quizy-be/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func setupAuthRoutes(r *chi.Mux) {
	client := db.Client
	if client == nil {
		panic("MongoDB client not initialized")
	}

	handlers.InitAuthHandlers(client.Database("quizy"))

	r.Post("/auth/signup", middleware.ErrorHandler(handlers.Signup))
}
