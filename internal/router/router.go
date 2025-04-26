package router

import "github.com/go-chi/chi/v5"

func SetupRouter(r *chi.Mux) {
	setupAuthRoutes(r)
}
