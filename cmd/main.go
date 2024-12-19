package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/cmd/dependencies"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
)

func main() {
	// db := database.CreateDatabase()

	// // repositories setup
	// employeeRp := repository.CreateEmployeeRepository(db.TbEmployees)

	// // services
	// employeeSv := service.CreateEmployeeService(employeeRp)

	// // handlers
	// employeeHd := handler.CreateEmployeeHandler(employeeSv)

	productHandler, employeeHd, sellersHandler, buyerHandler := dependencies.LoadDependencies()

	// routes setup
	rt := initRoutes(productHandler, employeeHd, sellersHandler, buyerHandler)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(productHandler *handler.ProductHandler,
	employeeHd *handler.EmployeeHandler, sellersHandler *handler.SellersController,
	buyerHandler *handler.BuyerHandler) *chi.Mux {
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
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		r.Get("/", buyerHandler.HandlerGetAllBuyers)
		r.Get("/{id}", buyerHandler.HandlerGetBuyerById)
		r.Post("/", buyerHandler.HandlerCreateBuyer)
		r.Patch("/{id}", buyerHandler.HandlerUpdateBuyer)
		r.Delete("/{id}", buyerHandler.HandlerDeleteBuyerById)
	})

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		r.Get("/", nil)
		r.Get("/{id}", nil)
		r.Post("/", nil)
		r.Patch("/{id}", nil)
		r.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", employeeHd.GetEmployeesHandler)
		// rt.Get("/{id}", nil)
		// rt.Post("/", nil)
		// rt.Patch("/{id}", nil)
		// rt.Delete("/{id}", nil)
	})
	return rt
}
