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
	localitiesService := service.CreateServiceLocalities(localitiesRepository, logInstance)
	localitiesHandler := handler.CreateHandlerLocality(localitiesService, logInstance)

	sellersRepository := repository.CreateRepositorySellers(sqlDB, logInstance)
	sellersService := service.CreateServiceSellers(sellersRepository, localitiesService, logInstance)
	sellersHandler := handler.CreateHandlerSellers(sellersService, logInstance)

	productRepo := repository.NewProductRepository(sqlDB, logInstance)
	productServ := service.NewProductService(productRepo, sellersRepository, logInstance)
	productHandler := handler.NewProductHandler(productServ, logInstance)

	productRecordRepo := repository.NewProductRecRepository(sqlDB, logInstance)
	productRecordServ := service.NewProductRecService(productRecordRepo, productServ, logInstance)
	productRecordHandler := handler.NewProductRecHandler(productRecordServ, logInstance)

	buyersRepository := repository.NewBuyerRepository(sqlDB, logInstance)
	buyerService := service.NewBuyerService(buyersRepository, logInstance)
	buyerHandler := handler.NewBuyerHandler(buyerService, logInstance)

	warehousesRepository := repository.NewWareHouseRepository(sqlDB, logInstance)
	warehousesService := service.NewWareHouseService(warehousesRepository, logInstance)
	warehousesHandler := handler.NewWareHouseHandler(warehousesService, logInstance)

	sectionsRep := repository.CreateRepositorySections(sqlDB, logInstance)
	sectionsSvc := service.CreateServiceSection(sectionsRep, logInstance)
	sectionsHandler := handler.CreateHandlerSections(sectionsSvc, logInstance)

	employeeRp := repository.CreateEmployeeRepository(sqlDB, logInstance)
	employeeSv := service.CreateEmployeeService(employeeRp, warehousesRepository, logInstance)
	employeeHd := handler.CreateEmployeeHandler(employeeSv, logInstance)

	inboundRp := repository.NewInboundService(sqlDB, logInstance)
	inboundSv := service.NewInboundOrderService(inboundRp, employeeSv, warehousesService, logInstance)
	inboundHd := handler.NewInboundHandler(inboundSv, logInstance)

	purchaseOrderRepository := repository.NewPurchaseOrderRepository(sqlDB, logInstance)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepository, buyerService, productRecordServ, logInstance)
	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService, logInstance)

	productBatchesRep := repository.CreateProductBatchesRepository(sqlDB, logInstance)
	productBatchesSvc := service.CreateProductBatchesService(productBatchesRep, productServ, sectionsSvc, logInstance)
	productBatchesHandler := handler.CreateProductBatchesHandler(productBatchesSvc, logInstance)

	carrierRep := repository.NewCarriersRepository(sqlDB, logInstance)
	carrierSv := service.NewCarrierService(carrierRep, localitiesService, logInstance)
	carrierHd := handler.NewCarrierHandler(carrierSv, logInstance)

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionsHandler, purchaseOrderHandler, inboundHd, productRecordHandler, productBatchesHandler, localitiesHandler, carrierHd
}
