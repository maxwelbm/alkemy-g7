package cmd

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {

	rt := initRoutes()

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes() *chi.Mux {
	rt := chi.NewRouter()

	rt.Route("/api/v1/warehouses", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sections", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/products", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	return rt
}
