package dependencies

import (
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
	"github.com/maxwelbm/alkemy-g7.git/pkg/database"
)

func LoadDependencies() (*handler.ProductHandler, *handler.EmployeeHandler, *handler.SellersController, *handler.BuyerHandler, *handler.WarehouseHandler, *handler.SectionController) {
	db := database.CreateDatabase()

	sellersRepository := repository.CreateRepositorySellers(db.TbSellers)
	sellersService := service.CreateServiceSellers(sellersRepository)
	sellersHandler := handler.CreateHandlerSellers(sellersService)

	productRepo := repository.NewProductRepository(*db)
	productServ := service.NewProductService(productRepo, sellersRepository)
	productHandler := handler.NewProductHandler(productServ)

	buyersRepository := repository.NewBuyerRepository(*db)
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

	return productHandler, employeeHd, sellersHandler, buyerHandler, warehousesHandler, sectionsHandler
}
