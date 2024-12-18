package dependencies

import (
	"github.com/maxwelbm/alkemy-g7.git/cmd/database"
	"github.com/maxwelbm/alkemy-g7.git/internal/handler"
	"github.com/maxwelbm/alkemy-g7.git/internal/repository"
	"github.com/maxwelbm/alkemy-g7.git/internal/service"
)

func LoadDependencies() *handler.ProductHandler {
	db := database.CreateDatabase()

	productRepo := repository.NewProductRepository(*db)
	productServ := service.NewProductService(productRepo)
	productHandler := handler.NewProductHandler(productServ)

	return productHandler
}