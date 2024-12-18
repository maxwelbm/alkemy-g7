package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

func main() {

	dbBuyer := database.CreateDatabase()
	rpBuyer := repository.NewBuyerRepository(*dbBuyer)
	svcBuyer := service.NewBuyerService(*rpBuyer)
	hdBuyer := handler.NewBuyerHandler(svcBuyer) // Passando o ponteiro para o handler
	rt := initRoutes(hdBuyer)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(hdBuyer *handler.BuyerHandler) *chi.Mux {
	rt := chi.NewRouter()

	// rt.Route("/api/v1/warehouses", func(r chi.Router) {
	// 	rt.Get("/", nil)
	// 	rt.Get("/{id}", nil)
	// 	rt.Post("/", nil)
	// 	rt.Patch("/{id}", nil)
	// 	rt.Delete("/{id}", nil)
	// })

	// rt.Route("/api/v1/sections", func(r chi.Router) {
	// 	rt.Get("/", nil)
	// 	rt.Get("/{id}", nil)
	// 	rt.Post("/", nil)
	// 	rt.Patch("/{id}", nil)
	// 	rt.Delete("/{id}", nil)
	// })

	// rt.Route("/api/v1/products", func(r chi.Router) {
	// 	rt.Get("/", nil)
	// 	rt.Get("/{id}", nil)
	// 	rt.Post("/", nil)
	// 	rt.Patch("/{id}", nil)
	// 	rt.Delete("/{id}", nil)
	// })

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		r.Get("/", hdBuyer.HandlerGetAllBuyers)
		// rt.Get("/{id}", nil)
		// rt.Post("/", nil)
		// rt.Patch("/{id}", nil)
		// rt.Delete("/{id}", nil)
	})

	// rt.Route("/api/v1/sellers", func(r chi.Router) {
	// 	rt.Get("/", nil)
	// 	rt.Get("/{id}", nil)
	// 	rt.Post("/", nil)
	// 	rt.Patch("/{id}", nil)
	// 	rt.Delete("/{id}", nil)
	// })

	return rt
}
