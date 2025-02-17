package dependencies

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

func LoadDependencies(sqlDB *sql.DB, logInstance logger.Logger) (*handler.ProductHandler, *handler.EmployeeHandler,
	*handler.SellersController, *handler.BuyerHandler, *handler.WarehouseHandler,
	*handler.SectionController, *handler.PurchaseOrderHandler, *handler.InboundOrderHandler,
	*handler.ProductRecHandler, *handler.ProductBatchesController, *handler.LocalitiesController, *handler.CarrierHandler) {
	localitiesRepository := repository.CreateRepositoryLocalities(sqlDB, logInstance)
	localitiesService := service.CreateServiceLocalities(localitiesRepository)
	localitiesHandler := handler.CreateHandlerLocality(localitiesService)

	sellersRepository := repository.CreateRepositorySellers(sqlDB, logInstance)
	sellersService := service.CreateServiceSellers(sellersRepository, localitiesService)
	sellersHandler := handler.CreateHandlerSellers(sellersService)

	productRepo := repository.NewProductRepository(sqlDB, logInstance)
	productServ := service.NewProductService(productRepo, sellersRepository)
	productHandler := handler.NewProductHandler(productServ)

	productRecordRepo := repository.NewProductRecRepository(sqlDB, logInstance)
	productRecordServ := service.NewProductRecService(productRecordRepo, productServ)
	productRecordHandler := handler.NewProductRecHandler(productRecordServ)

	buyersRepository := repository.NewBuyerRepository(sqlDB, logInstance)
	buyerService := service.NewBuyerService(buyersRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	warehousesRepository := repository.NewWareHouseRepository(sqlDB, logInstance)
	warehousesService := service.NewWareHouseService(warehousesRepository)
	warehousesHandler := handler.NewWareHouseHandler(warehousesService)

	sectionsRep := repository.CreateRepositorySections(sqlDB, logInstance)
	sectionsSvc := service.CreateServiceSection(sectionsRep)
	sectionsHandler := handler.CreateHandlerSections(sectionsSvc)

	employeeRp := repository.CreateEmployeeRepository(sqlDB, logInstance)
	employeeSv := service.CreateEmployeeService(employeeRp, warehousesRepository)
	employeeHd := handler.CreateEmployeeHandler(employeeSv)

	inboundRp := repository.NewInboundService(sqlDB, logInstance)
	inboundSv := service.NewInboundOrderService(inboundRp, employeeSv, warehousesService)
	inboundHd := handler.NewInboundHandler(inboundSv)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository(sqlDB, logInstance)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepository, buyerService, productRecordServ)
	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	productBatchesRep := repository.CreateProductBatchesRepository(sqlDB, logInstance)
	productBatchesSvc := service.CreateProductBatchesService(productBatchesRep, productServ, sectionsSvc)
	productBatchesHandler := handler.CreateProductBatchesHandler(productBatchesSvc)

	carrierRep := repository.NewCarriersRepository(sqlDB, logInstance)
	carrierSv := service.NewCarrierService(carrierRep, localitiesService)
	carrierHd := handler.NewCarrierHandler(carrierSv)

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionsHandler, purchaseOrderHandler, inboundHd, productRecordHandler, productBatchesHandler, localitiesHandler, carrierHd
}
