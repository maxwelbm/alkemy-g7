package repository

import (
	"database/sql"
	"fmt"

	"github.com/maxwelbm/alkemy-g7.git/pkg/logger"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductBatchesRepository struct {
	db  *sql.DB
	log logger.Logger
}

func CreateProductBatchesRepository(db *sql.DB, log logger.Logger) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db, log: log}
}

func (r *ProductBatchesRepository) GetByID(id int) (prodBatches model.ProductBatches, err error) {
	r.log.Log("ProductBatchesRepository", "INFO", "initializing GetByID function")

	getByIDQuery := "SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?"
	row := r.db.QueryRow(getByIDQuery, id)

	err = row.Scan(&prodBatches.ID, &prodBatches.BatchNumber, &prodBatches.CurrentQuantity, &prodBatches.CurrentTemperature, &prodBatches.MinimumTemperature, &prodBatches.DueDate, &prodBatches.InitialQuantity, &prodBatches.ManufacturingDate, &prodBatches.ManufacturingHour, &prodBatches.ProductID, &prodBatches.SectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("product batches", customerror.ErrorNotFound, "")
		}

		r.log.Log("ProductBatchesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	r.log.Log("ProductBatchesRepository", "INFO", fmt.Sprintf("returning a product batches by id with no error: %v", prodBatches))

	return
}

func (r *ProductBatchesRepository) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	r.log.Log("ProductBatchesRepository", "INFO", "initializing Post function")

	postQuery := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, minimum_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(postQuery, (*prodBatches).BatchNumber, (*prodBatches).CurrentQuantity, (*prodBatches).CurrentTemperature, (*prodBatches).MinimumTemperature, (*prodBatches).DueDate, (*prodBatches).InitialQuantity, (*prodBatches).ManufacturingDate, (*prodBatches).ManufacturingHour, (*prodBatches).ProductID, (*prodBatches).SectionID)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("product batches:", customerror.ErrorConflict, "")
		}

		r.log.Log("ProductBatchesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		r.log.Log("ProductBatchesRepository", "ERROR", fmt.Sprintf("Error: %v", err))

		return
	}

	(*prodBatches).ID = int(id)

	newProdBatches, _ = r.GetByID(int(id))

	r.log.Log("ProductBatchesRepository", "INFO", fmt.Sprintf("saving a product batches to the database: %v", prodBatches))

	return
}
