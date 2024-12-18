package dependencies

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

func LoadDependencies() (*handler.ProductHandler, *handler.EmployeeHandler, *handler.SellersController, *handler.BuyerHandler, *handler.WarehouseHandler) {
	db := database.CreateDatabase()

	employeeRp := repository.CreateEmployeeRepository(db.TbEmployees)
	employeeSv := service.CreateEmployeeService(employeeRp)
	employeeHd := handler.CreateEmployeeHandler(employeeSv)

	productRepo := repository.NewProductRepository(*db)
	productServ := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productServ)

	sellersRepository := repository.CreateRepositorySellers(db.TbSellers)
	sellersService := service.CreateServiceSellers(*sellersRepository)
	sellersHandler := handler.CreateHandlerSellers(*sellersService)

	buyersRepository := repository.NewBuyerRepository(*db)
	buyerService := service.NewBuyerService(buyersRepository)
	buyerHandler := handler.NewBuyerHandler(buyerService)

	warehousesRepository := repository.NewWareHouseRepository(*db)
	warehousesService := service.NewWareHoureService(*warehousesRepository)
	warehousesHandler := handler.NewWareHouseHandler(*warehousesService)

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler
}
