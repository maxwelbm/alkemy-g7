package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/customerror"
)

type ProductBatchesRepository struct {
	db *sql.DB
}

func CreateProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db}
}

func (r *ProductBatchesRepository) GetByID(id int) (prodBatches model.ProductBatches, err error) {
	getByIDQuery := "SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?"
	row := r.db.QueryRow(getByIDQuery, id)

	err = row.Scan(&prodBatches.ID, &prodBatches.BatchNumber, &prodBatches.CurrentQuantity, &prodBatches.CurrentTemperature, &prodBatches.MinimumTemperature, &prodBatches.DueDate, &prodBatches.InitialQuantity, &prodBatches.ManufacturingDate, &prodBatches.ManufacturingHour, &prodBatches.ProductID, &prodBatches.SectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = customerror.HandleError("product batches", customerror.ErrorNotFound, "")
		}

		return
	}

	return
}

func (r *ProductBatchesRepository) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	postQuery := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, minimum_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(postQuery, (*prodBatches).BatchNumber, (*prodBatches).CurrentQuantity, (*prodBatches).CurrentTemperature, (*prodBatches).MinimumTemperature, (*prodBatches).DueDate, (*prodBatches).InitialQuantity, (*prodBatches).ManufacturingDate, (*prodBatches).ManufacturingHour, (*prodBatches).ProductID, (*prodBatches).SectionID)

	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = customerror.HandleError("product batches:", customerror.ErrorConflict, "")
		}
		
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*prodBatches).ID = int(id)

	newProdBatches, _ = r.GetByID(int(id))

	return
}
