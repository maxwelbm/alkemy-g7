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
	db := database.CreateDatabase()

	// repositories setup
	employeeRp := repository.CreateEmployeeRepository(db.TbEmployees)

	// services
	employeeSv := service.CreateEmployeeService(employeeRp)

	// handlers
	employeeHd := handler.CreateEmployeeHandler(employeeSv)

	// routes setup
	rt := initRoutes(employeeHd)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}

func initRoutes(employeeHandler *handler.EmployeeHandler) *chi.Mux {
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

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		rt.Get("/", nil)
		rt.Get("/{id}", nil)
		rt.Post("/", nil)
		rt.Patch("/{id}", nil)
		rt.Delete("/{id}", nil)
	})

	rt.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", employeeHandler.GetEmployeesHandler)
		r.Get("/{id}", employeeHandler.GetEmployeeById)
		// rt.Post("/", nil)
		// rt.Patch("/{id}", nil)
		// rt.Delete("/{id}", nil)
	})
	return rt
}
