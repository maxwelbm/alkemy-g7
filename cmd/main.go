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
	dbWareHouse := database.CreateDatabase()
	rpWareHouse := repository.NewWareHouseMap(*dbWareHouse)
	srvWareHouse := service.NewWareHouDefault(*rpWareHouse)
	hdWareHouse := handler.NewWareHouseHandlerDefault(*srvWareHouse)

	rt := initRoutes(hdWareHouse)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(hd *handler.WarehouseHandler) *chi.Mux {
	rt := chi.NewRouter()

	rt.Route("/api/v1/warehouses", func(r chi.Router) {
		r.Get("/", hd.GetAllWareHouse())
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
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

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	return rt
}
