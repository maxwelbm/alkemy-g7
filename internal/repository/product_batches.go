package repository

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
	"github.com/maxwelbm/alkemy-g7.git/internal/model"
	"github.com/maxwelbm/alkemy-g7.git/pkg/custom_error"
)

type ProductBatchesRepository struct {
	db *sql.DB
}

func CreateProductBatchesRepository(db *sql.DB) *ProductBatchesRepository {
	return &ProductBatchesRepository{db: db}
}

func (r *ProductBatchesRepository) GetById(id int) (prodBatches model.ProductBatches, err error) {
	getByIdQuery := "SELECT `id`, `batch_number`, `current_quantity`, `current_temperature`, `minimum_temperature`, `due_date`, `initial_quantity`, `manufacturing_date`, `manufacturing_hour`, `product_id`, `section_id` FROM `product_batches` WHERE `id` = ?"
	row := r.db.QueryRow(getByIdQuery, id)

	err = row.Scan(&prodBatches.ID, &prodBatches.BatchNumber, &prodBatches.CurrentQuantity, &prodBatches.CurrentTemperature, &prodBatches.MinimumTeperature, &prodBatches.DueDate, &prodBatches.InitialQuantity, &prodBatches.ManufacturingDate, &prodBatches.ManufacturingHour, &prodBatches.ProductID, &prodBatches.SectionID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = custom_error.NotFound
		}
		return
	}
	return
}

func (r *ProductBatchesRepository) Post(prodBatches *model.ProductBatches) (newProdBatches model.ProductBatches, err error) {
	postQuery := "INSERT INTO product_batches (batch_number, current_quantity, current_temperature, minimum_temperature, due_date, initial_quantity, manufacturing_date, manufacturing_hour, product_id, section_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(postQuery, (*prodBatches).BatchNumber, (*prodBatches).CurrentQuantity, (*prodBatches).CurrentTemperature, (*prodBatches).MinimumTeperature, (*prodBatches).DueDate, (*prodBatches).InitialQuantity, (*prodBatches).ManufacturingDate, (*prodBatches).ManufacturingHour, (*prodBatches).ProductID, (*prodBatches).SectionID)
	if err != nil {
		if err.(*mysql.MySQLError).Number == 1062 {
			err = custom_error.Conflict
		}
		return
	}

	id, err := result.LastInsertId()
	if err != nil {
		return
	}

	(*prodBatches).ID = int(id)

	newProdBatches, _ = r.GetById(int(id))

	return
}
