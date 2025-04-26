package main

import (
	"net/http"
	middleware "quizy-be/internal"
	"quizy-be/internal/router"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	middleware.SetupMiddleware(r)
	router.SetupRouter(r)

	http.ListenAndServe(":8080", r)
}
