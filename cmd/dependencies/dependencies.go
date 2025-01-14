package dependencies

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

func LoadDependencies(slqDb *sql.DB) (*handler.ProductHandler, *handler.EmployeeHandler,
	*handler.SellersController, *handler.BuyerHandler, *handler.WarehouseHandler,
	*handler.SectionController, *handler.PurchaseOrderHandler, *handler.InboundOrderHandler,
	*handler.ProductRecHandler, *handler.ProductBatchesController) {

	sellersRepository := repository.CreateRepositorySellers(slqDb)
	sellersService := service.CreateServiceSellers(sellersRepository)
	sellersHandler := handler.CreateHandlerSellers(sellersService)

	productRepo := repository.NewProductRepository(slqDb)
	productServ := service.NewProductService(productRepo, sellersRepository)
	productHandler := handler.NewProductHandler(productServ)

	productRecordRepo := repository.NewProductRecRepository(slqDb)
	productRecordServ := service.NewProductRecService(productRecordRepo, productServ)
	productRecordHandler := handler.NewProductRecHandler(productRecordServ)

	buyersRepository := repository.NewBuyerRepository(slqDb)
	buyerService := service.NewBuyerService(buyersRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	warehousesRepository := repository.NewWareHouseRepository(slqDb)
	warehousesService := service.NewWareHoureService(warehousesRepository)
	warehousesHandler := handler.NewWareHouseHandler(warehousesService)

	sectionsRep := repository.CreateRepositorySections(slqDb)
	sectionsSvc := service.CreateServiceSection(*sectionsRep)
	sectionsHandler := handler.CreateHandlerSections(sectionsSvc)

	employeeRp := repository.CreateEmployeeRepository(slqDb)
	employeeSv := service.CreateEmployeeService(employeeRp, warehousesRepository)
	employeeHd := handler.CreateEmployeeHandler(employeeSv)

	inboundRp := repository.NewInboundService(slqDb)
	inboundSv := service.NewInboundOrderService(inboundRp, employeeSv, warehousesService)
	inboundHd := handler.NewInboundHandler(inboundSv)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository(slqDb)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepository, buyerService, productRecordServ)
	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	productBatchesRep := repository.CreateProductBatchesRepository(slqDb)
	productBatchesSvc := service.CreateProductBatchesService(*productBatchesRep, productServ, sectionsSvc)
	productBatchesHandler := handler.CreateProductBatchesHandler(productBatchesSvc)

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionsHandler, purchaseOrderHandler, inboundHd, productRecordHandler, productBatchesHandler

}
