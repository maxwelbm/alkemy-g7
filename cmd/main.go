package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/maxwelbm/alkemy-g7.git/cmd/dependencies"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

func main() {
	dbConfig, err := database.GetDbConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnectionDb(dbConfig)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionHandler := dependencies.LoadDependencies(db.Connection)

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
		r.Post("/", warehouseHandler.PostWareHouse())
		r.Patch("/{id}", warehouseHandler.UpdateWareHouse())
		r.Delete("/{id}", warehouseHandler.DeleteByIdWareHouse())
	})

	rt.Route("/api/v1/sections", func(r chi.Router) {
		r.Get("/", sectionHandler.GetAll)
		r.Get("/{id}", sectionHandler.GetById)
		r.Post("/", sectionHandler.Post)
		r.Patch("/{id}", sectionHandler.Update)
		r.Delete("/{id}", sectionHandler.Delete)
		r.Get("/reportProducts", nil)
	})

	rt.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", productHandler.GetAllProducts)
		r.Get("/{id}", productHandler.GetProductById)
		r.Post("/", productHandler.CreateProduct)
		r.Patch("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProductById)
		r.Get("/reportRecords", nil)
	})

	rt.Route("/api/v1/buyers", func(r chi.Router) {
		r.Get("/", buyerHandler.HandlerGetAllBuyers)
		r.Get("/{id}", buyerHandler.HandlerGetBuyerById)
		r.Post("/", buyerHandler.HandlerCreateBuyer)
		r.Patch("/{id}", buyerHandler.HandlerUpdateBuyer)
		r.Delete("/{id}", buyerHandler.HandlerDeleteBuyerById)
		r.Get("/reportPurchaseOrders", nil)
	})

	rt.Route("/api/v1/sellers", func(r chi.Router) {
		r.Get("/", sellersHandler.GetAllSellers)
		r.Get("/{id}", sellersHandler.GetById)
		r.Post("/", sellersHandler.CreateSellers)
		r.Patch("/{id}", sellersHandler.UpdateSellers)
		r.Delete("/{id}", sellersHandler.DeleteSellers)
	})

	rt.Route("/api/v1/employees", func(r chi.Router) {
		r.Get("/", employeeHd.GetEmployeesHandler)
		r.Get("/{id}", employeeHd.GetEmployeeById)
		r.Post("/", employeeHd.InsertEmployee)
		r.Patch("/{id}", employeeHd.UpdateEmployee)
		r.Delete("/{id}", employeeHd.DeleteEmployee)
		r.Get("/reportInboundOrders", nil)
	})

	rt.Route("/api/v1/localities", func(r chi.Router) {
		r.Post("/", nil)
		r.Get("/reportCarries", nil)
		r.Get("/reportSellers", nil)
	})

	rt.Route("/api/v1/carries", func(r chi.Router) {
		r.Post("/", nil)
	})

	rt.Route("/api/v1/productBatches", func(r chi.Router) {
		r.Post("/", nil)
	})

	rt.Route("/api/v1/productRecords", func(r chi.Router) {
		r.Post("/", nil)
	})

	rt.Route("/api/v1/inboundOrders", func(r chi.Router) {
		r.Post("/", nil)
	})

	rt.Route("/api/v1/purchaseOrders", func(r chi.Router) {
		r.Post("/", nil)
	})

	return rt
}
