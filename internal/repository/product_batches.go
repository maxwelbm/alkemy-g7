package repository

import (
	"database/sql"

	"github.com/maxwelbm/alkemy-g7.git/internal/model"
)

type ProductBatchesRepository struct {
	db *sql.DB
}

func CreateProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db}
}

func (r *ProductBatchesRepository) GetById(id int) (prodBatches model.ProductBatches, err error) {
	return
}

func (r *ProductBatchesRepository) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	postQuery := "INSERT INTO `product_batches` (`batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(postQuery, (*prodBatches).BatchNumber, (*prodBatches).CurrentQuantity, (*prodBatches).CurrentTemperature, (*prodBatches).MinimumTeperature, (*prodBatches).DueDate, (*prodBatches).InitialQuantity, (*prodBatches).ManufacturingDate, (*prodBatches).ManufacturingHour, (*prodBatches).ProductID, (*prodBatches).SectionID)
	if err != nil {
		return
	}

	lastInsertId, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*prodBatches).ID = int(lastInsertId)

	return
}
