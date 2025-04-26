package main

import (
	"net/http"
	middleware "quizy-be/internal"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	middleware.SetupMiddleware(r)

	http.ListenAndServe(":8080", r)
}
