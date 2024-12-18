package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/cmd/database"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

func main() {

	db := database.CreateDatabase()
	repository := repository.CreateRepositorySellers(db.TbSellers)
	service := service.CreateServiceSellers(*repository)
	handler := handler.CreateHandlerSellers(*service)

	rt := initRoutes(*handler)
	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(handler handler.SellersController) *chi.Mux {
	rt := chi.NewRouter()

	rt.Route("/api/v1/warehouses", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sections", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/employee", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		r.Get("/", handler.GetAllSellers())
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	return rt
}
