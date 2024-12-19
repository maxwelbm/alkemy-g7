package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/cmd/dependencies"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
)

func main() {

	productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionHandler := dependencies.LoadDependencies()

	rt := initRoutes(productHandler, employeeHd, sellersHandler, buyerHandler, sectionHandler, warehousesHandler)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(productHandler *handler.ProductHandler,
	employeeHd *handler.EmployeeHandler, sellersHandler *handler.SellersController,
	buyerHandler *handler.BuyerHandler, sectionHandler *handler.SectionController, warehouseHandler *handler.WarehouseHandler) *chi.Mux {
	rt := chi.NewRouter()

	rt.Route("/api/v1/warehouses", func(r chi.Router) {
		r.Get("/", warehouseHandler.GetAllWareHouse())
		r.Get("/{id}", warehouseHandler.GetWareHouseById())
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sections", func(r chi.Router) {
		r.Get("/", sectionHandler.GetAll)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		r.Get("/", buyerHandler.HandlerGetAllBuyers)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		r.Get("/", sellersHandler.GetAllSellers)
		r.Get("/{id}", sellersHandler.GetById)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", employeeHd.GetEmployeesHandler)
		r.Get("/{id}", employeeHd.GetEmployeeById)
		r.Post("/", employeeHd.InsertEmployee)
		r.Patch("/{id}", employeeHd.UpdateEmployee)
		// rt.Delete("/{id}", nil)
	})
	return rt
}
