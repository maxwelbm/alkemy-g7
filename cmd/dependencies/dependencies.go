package dependencies

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

func LoadDependencies(slqDb *sql.DB) (*handler.ProductHandler, *handler.EmployeeHandler, *handler.SellersController, *handler.BuyerHandler, *handler.WarehouseHandler, *handler.SectionController, *handler.PurchaseOrderHandler) {
	db := database.CreateDatabase()

	sellersRepository := repository.CreateRepositorySellers(db.TbSellers)
	sellersService := service.CreateServiceSellers(sellersRepository)
	sellersHandler := handler.CreateHandlerSellers(sellersService)

	productRepo := repository.NewProductRepository(*db)
	productServ := service.NewProductService(productRepo, sellersRepository)
	productHandler := handler.NewProductHandler(productServ)

	buyersRepository := repository.NewBuyerRepository(slqDb)
	buyerService := service.NewBuyerService(buyersRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	warehousesRepository := repository.NewWareHouseRepository(*db)
	warehousesService := service.NewWareHoureService(warehousesRepository)
	warehousesHandler := handler.NewWareHouseHandler(warehousesService)

	sectionsRep := repository.CreateRepositorySections(*db)
	sectionsSvc := service.CreateServiceSection(*sectionsRep)
	sectionsHandler := handler.CreateHandlerSections(sectionsSvc)

	employeeRp := repository.CreateEmployeeRepository(db.TbEmployees)
	employeeSv := service.CreateEmployeeService(employeeRp, warehousesRepository)
	employeeHd := handler.CreateEmployeeHandler(employeeSv)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository(slqDb)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepository, buyerService, productRecordHandler)
	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionsHandler, purchaseOrderHandler
}
