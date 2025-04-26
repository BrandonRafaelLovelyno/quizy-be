package router

import (
	handler "quizy-be/internal/handlers"

	"github.com/go-chi/chi/v5"
)

func setupAuthRoutes(r *chi.Mux) {
	r.Post("/auth/login", handler.Login)
	r.Post("/auth/signup", handler.Signup)
}
